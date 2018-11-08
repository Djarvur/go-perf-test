#!/bin/sh -eux

go run twain/testTwain.go < mtent12.txt
perl   twain/testTwain.pl < mtent12.txt
go run email/testEMail.go < 100000USAEmailAddress.TXT
perl   email/testEMail.pl < 100000USAEmailAddress.TXT