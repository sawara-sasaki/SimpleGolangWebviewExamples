#!/bin/bash
BASE_DIR=`dirname $0`/../
TARGET_FILE=${BASE_DIR}src/static/view.tpl
TMP_FILE=${TARGET_FILE}.tmp

IMAGES="${BASE_DIR}src/img/*"
FILEARY=()
for FILEPATH in ${IMAGES}; do
  if [ -f ${FILEPATH} ] ; then
    FILEARY+=("${FILEPATH}")
  fi
done

for i in ${FILEARY[@]}; do
  FILENAME=`basename ${i}`
  if [ ${FILENAME##*.} == jpg ] ; then
    echo '{{define "'${FILENAME}'"}}data:image/jpeg;base64,' >> ${TMP_FILE}
    base64 -i ${i} >> ${TMP_FILE}
    echo '{{end}}' >> ${TMP_FILE}
  elif [ ${FILENAME##*.} == png ] ; then
    echo '{{define "'${FILENAME}'"}}data:image/png;base64,' >> ${TMP_FILE}
    base64 -i ${i} >> ${TMP_FILE}
    echo '{{end}}' >> ${TMP_FILE}
  fi
done

# base64url encode
cat ${TMP_FILE} | awk '{gsub("\\+", "-", $0);gsub("/", "_", $0);printf $0}' > ${TARGET_FILE}
rm ${TMP_FILE}
