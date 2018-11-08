use strict;

use Time::HiRes qw(time);

my $re = '\s*[a-zA-Z0-9.!#$%&\'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*\s+';

my $tString = '';

{
    local $/;
    $tString = <STDIN>;
}

printf("testing on %d bytes\n", length($tString));

my $sTime = time();
my @matches = $tString =~ m/$re/g;
my $i = scalar(@matches);

printf("perl regexp: %fs, %d found\n", time()-$sTime, $i);

