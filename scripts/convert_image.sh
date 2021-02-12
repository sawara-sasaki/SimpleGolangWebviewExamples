#!/bin/bash
BASE_DIR=`dirname $0`/../
TARGET_DIR="${BASE_DIR}static/base64/"

IMAGES="${BASE_DIR}static/img/*"
FILEARY=()
for FILEPATH in ${IMAGES}; do
  if [ -f ${FILEPATH} ] ; then
    FILEARY+=("${FILEPATH}")
  fi
done

TMP_FILE="${BASE_DIR}static/tmp"
for i in ${FILEARY[@]}; do
  FILENAME=`basename ${i}`
  TARGET_FILE=${TARGET_DIR}${FILENAME}.txt
  if [ ${FILENAME##*.} == jpg ] ; then
    echo 'data:image/jpeg;base64,' > ${TMP_FILE}
    base64 ${i} >> ${TMP_FILE}
  elif [ ${FILENAME##*.} == png ] ; then
    echo 'data:image/png;base64,' > ${TMP_FILE}
    base64 ${i} >> ${TMP_FILE}
  fi
  cat ${TMP_FILE} | awk '{printf $0}' > ${TARGET_FILE}
done
rm ${TMP_FILE}
