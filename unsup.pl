#!/usr/bin/perl
# paul@unsup.sbrk.co.uk
use Compress::Zlib;
use strict;
 
my $inf = shift;
my $in;
open(F, "<$inf") || die "Usage: unsup.pl /path/to/supout.inf";
while (<F>) {
        $_ =~ s/\s+$//; # strip terminating \r \n or other whitespace chars
        if ($_ eq "--BEGIN ROUTEROS SUPOUT SECTION") {
                $in = "";
        } elsif ($_ eq "--END ROUTEROS SUPOUT SECTION") {
                decode($in);
        } else {
                $in .= $_;
        }
}
 
# this is base64 but done in a different byte order
sub decode {
        my $in = shift;
        # terminating "=" is so that index %64 == 0 for pad char
        my $b64 = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";
        #my $np = length($in)-index($in,"="); # ignored at the moment
        #$np = 0 if (-1 == index($in,"="));
        my $out;
        for (my $i = 0; $i < length($in); $i+=4) {
                my $o = index($b64, substr($in,$i+3,1))%64 << 18 |
                        index($b64, substr($in,$i+2,1))%64 << 12 |
                        index($b64, substr($in,$i+1,1))%64 << 6 |
                        index($b64, substr($in,$i,1))%64;
                $out .= chr($o%256);
                $out .= chr(($o>>8) % 256);
                $out .= chr(($o>>16) % 256);
        }
        # decoded data consists of "section_name\0zlib_compressed_data"
        my $sec = substr($out, 0, index($out,"\0"));
        print "==SECTION $sec\n";
        my $cmp = substr($out, index($out,"\0")+1);
        my $uncomp = uncompress($cmp);
        print "$uncomp\n";
}