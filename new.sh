#!/bin/bash

MKDIR=$(which mkdir)
# echo $MKDIR
TOUCH=$(which touch)
# echo $TOUCH
DIR=$1
$MKDIR $DIR
MAINCODE="${DIR}/main.go"
$TOUCH $MAINCODE

echo "package main" > $MAINCODE
echo "import (" >> $MAINCODE
echo "    \"fmt\"" >> $MAINCODE
echo "    \"os\"" >> $MAINCODE
echo ")" >> $MAINCODE
echo "" >> $MAINCODE

echo "func main() {" >> $MAINCODE
echo "" >> $MAINCODE
echo "}" >> $MAINCODE

exit 0