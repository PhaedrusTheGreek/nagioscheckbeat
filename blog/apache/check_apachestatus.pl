#!/usr/bin/perl -w
####################### check_apachestatus.pl #######################
# Version : 1.1
# Date : 27 Jul 2007 
# Author  : De Bodt Lieven (Lieven.DeBodt at gmail.com)
# Licence : GPL - http://www.fsf.org/licenses/gpl.txt
#############################################################
#
# help : ./check_apachestatus.pl -h

use strict;
use Getopt::Long;
use LWP::UserAgent;
use Time::HiRes qw(gettimeofday tv_interval);

# Nagios specific

use lib "/usr/local/nagios/libexec";
use utils qw(%ERRORS $TIMEOUT);
#my %ERRORS=('OK'=>0,'WARNING'=>1,'CRITICAL'=>2,'UNKNOWN'=>3,'DEPENDENT'=>4);

# Globals

my $Version='1.0';
my $Name=$0;

my $o_host =		undef; 		# hostname 
my $o_help=		undef; 		# want some help ?
my $o_port = 		undef; 		# port
my $o_version= 		undef;  	# print version
my $o_warn_level=	undef;  	# Number of available slots that will cause a warning
my $o_crit_level=	undef;  	# Number of available slots that will cause an error
my $o_timeout=  	15;            	# Default 15s Timeout

# functions

sub show_versioninfo { print "$Name version : $Version\n"; }

sub print_usage {
  print "Usage: $Name -H <host> [-p <port>] [-t <timeout>] [-w <warn_level> -c <crit_level>] [-V]\n";
}

# Get the alarm signal
$SIG{'ALRM'} = sub {
  print ("ERROR: Alarm signal (Nagios time-out)\n");
  exit $ERRORS{"CRITICAL"};
};

sub help {
  print "Apache Monitor for Nagios version ",$Version,"\n";
  print "GPL licence, (c)2006-2007 De Bodt Lieven\n\n";
  print_usage();
  print <<EOT;
-h, --help
   print this help message
-H, --hostname=HOST
   name or IP address of host to check
-p, --port=PORT
   Http port
-t, --timeout=INTEGER
   timeout in seconds (Default: $o_timeout)
-w, --warn=MIN
   number of available slots that will cause a warning
   -1 for no warning
-c, --critical=MIN
   number of available slots that will cause an error
-V, --version
   prints version number
Note :
  The script will return
    * Without warn and critical options:
        OK       if we are able to connect to the apache server's status page,
        CRITICAL if we aren't able to connect to the apache server's status page,,
    * With warn and critical options:
        OK       if we are able to connect to the apache server's status page and #available slots > <warn_level>,
        WARNING  if we are able to connect to the apache server's status page and #available slots <= <warn_level>,
        CRITICAL if we are able to connect to the apache server's status page and #available slots <= <crit_level>,
        UNKNOWN  if we aren't able to connect to the apache server's status page

Perfdata legend:
"_;S;R;W;K;D;C;L;G;I;."
_ : Waiting for Connection
S : Starting up
R : Reading Request
W : Sending Reply
K : Keepalive (read)
D : DNS Lookup
C : Closing connection
L : Logging
G : Gracefully finishing
I : Idle cleanup of worker
. : Open slot with no current process

EOT
}

sub check_options {
  Getopt::Long::Configure ("bundling");
  GetOptions(
      'h'     => \$o_help,        'help'          => \$o_help,
      'H:s'   => \$o_host,        'hostname:s'	  => \$o_host,
      'p:i'   => \$o_port,        'port:i'	  => \$o_port,
      'V'     => \$o_version,     'version'       => \$o_version,
      'w:i'   => \$o_warn_level,  'warn:i'	  => \$o_warn_level,
      'c:i'   => \$o_crit_level,  'critical:i'	  => \$o_crit_level,
      't:i'   => \$o_timeout,     'timeout:i'     => \$o_timeout,

  );

  if (defined ($o_help)) { help(); exit $ERRORS{"UNKNOWN"}};
  if (defined($o_version)) { show_versioninfo(); exit $ERRORS{"UNKNOWN"}};
  if (((defined($o_warn_level) && !defined($o_crit_level)) || (!defined($o_warn_level) && defined($o_crit_level))) || ((defined($o_warn_level) && defined($o_crit_level)) && (($o_warn_level != -1) &&  ($o_warn_level <= $o_crit_level)))) { 
    print "Check warn and crit!\n"; print_usage(); exit $ERRORS{"UNKNOWN"}
  }
  # Check compulsory attributes
  if (!defined($o_host)) { print_usage(); exit $ERRORS{"UNKNOWN"}};
}

########## MAIN ##########

check_options();

my $ua = LWP::UserAgent->new( protocols_allowed => ['http'], timeout => $o_timeout);
my $timing0 = [gettimeofday];
my $response = undef;
if (!defined($o_port)) {
  $response = $ua->get('http://' . $o_host . '/server-status');
} else {
  $response = $ua->get('http://' . $o_host . ':' . $o_port . '/server-status');
}
my $timeelapsed = tv_interval ($timing0, [gettimeofday]);

my $webcontent = undef;
if ($response->is_success) {
  $webcontent=$response->content;
  my @webcontentarr = split("\n", $webcontent);
  my $i = 0;
  my $BusyWorkers=undef;
  my $IdleWorkers=undef;
  # Get the amount of idle and busy workers(Apache2)/servers(Apache1)
  while (($i < @webcontentarr) && ((!defined($BusyWorkers)) || (!defined($IdleWorkers)))) {
    if ($webcontentarr[$i] =~ /(\d+)\s+requests\s+currently\s+being\s+processed,\s+(\d+)\s+idle\s+....ers/) {
      ($BusyWorkers, $IdleWorkers) = ($webcontentarr[$i] =~ /(\d+)\s+requests\s+currently\s+being\s+processed,\s+(\d+)\s+idle\s+....ers/);
    }
    $i++;
  }

  # Get the scoreboard
  my $ScoreBoard = "";
  $i = 0;
  my $PosPreBegin = undef;
  my $PosPreEnd = undef;
  while (($i < @webcontentarr) && ((!defined($PosPreBegin)) || (!defined($PosPreEnd)))) {
    if (!defined($PosPreBegin)) {
      if ( $webcontentarr[$i] =~ m/<pre>/i ) {
        $PosPreBegin = $i;
      }
    } 
    if (defined($PosPreBegin)) {
      if ( $webcontentarr[$i] =~ m/<\/pre>/i ) {
        $PosPreEnd = $i;
      }
    }
    $i++;
  }
  for ($i = $PosPreBegin; $i <= $PosPreEnd; $i++) {
    $ScoreBoard = $ScoreBoard . $webcontentarr[$i];
  }
  $ScoreBoard =~ s/^.*<[Pp][Rr][Ee]>//;
  $ScoreBoard =~ s/<\/[Pp][Rr][Ee].*>//;

  my $CountOpenSlots = ($ScoreBoard =~ tr/\.//);
  if (defined($o_crit_level) && ($o_crit_level != -1)) {
    if (($CountOpenSlots + $IdleWorkers) <= $o_crit_level) {
      printf("CRITICAL %f seconds response time. Idle %d, busy %d, open slots %d | %d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d\n", $timeelapsed, $IdleWorkers, $BusyWorkers, $CountOpenSlots, ($ScoreBoard =~ tr/\_//), ($ScoreBoard =~ tr/S//),($ScoreBoard =~ tr/R//),($ScoreBoard =~ tr/W//),($ScoreBoard =~ tr/K//),($ScoreBoard =~ tr/D//),($ScoreBoard =~ tr/C//),($ScoreBoard =~ tr/L//),($ScoreBoard =~ tr/G//),($ScoreBoard =~ tr/I//), $CountOpenSlots);
      exit $ERRORS{"CRITICAL"}
    }
  } 
  if (defined($o_warn_level) && ($o_warn_level != -1)) {
    if (($CountOpenSlots + $IdleWorkers) <= $o_warn_level) {
      printf("WARNING %f seconds response time. Idle %d, busy %d, open slots %d | %d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d\n", $timeelapsed, $IdleWorkers, $BusyWorkers, $CountOpenSlots, ($ScoreBoard =~ tr/\_//), ($ScoreBoard =~ tr/S//),($ScoreBoard =~ tr/R//),($ScoreBoard =~ tr/W//),($ScoreBoard =~ tr/K//),($ScoreBoard =~ tr/D//),($ScoreBoard =~ tr/C//),($ScoreBoard =~ tr/L//),($ScoreBoard =~ tr/G//),($ScoreBoard =~ tr/I//), $CountOpenSlots);
      exit $ERRORS{"WARNING"}
    }
  }
  printf("OK %f seconds response time. Idle %d, busy %d, open slots %d | %d;%d;%d;%d;%d;%d;%d;%d;%d;%d;%d\n", $timeelapsed, $IdleWorkers, $BusyWorkers, $CountOpenSlots, ($ScoreBoard =~ tr/\_//), ($ScoreBoard =~ tr/S//),($ScoreBoard =~ tr/R//),($ScoreBoard =~ tr/W//),($ScoreBoard =~ tr/K//),($ScoreBoard =~ tr/D//),($ScoreBoard =~ tr/C//),($ScoreBoard =~ tr/L//),($ScoreBoard =~ tr/G//),($ScoreBoard =~ tr/I//), $CountOpenSlots);
      exit $ERRORS{"OK"}
}
else {
  if (defined($o_warn_level) || defined($o_crit_level)) {
    printf("UNKNOWN %s\n", $response->status_line);
    exit $ERRORS{"UNKNOWN"}
  } else {
    printf("CRITICAL %s\n", $response->status_line);
    exit $ERRORS{"CRITICAL"}
  }
}
