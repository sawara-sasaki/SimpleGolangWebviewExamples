#!/bin/bash
BASE_DIR=`dirname $0`/../

# img.tpl
TARGET_FILE=${BASE_DIR}src/static/img.tpl
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

# css.tpl
TARGET_FILE=${BASE_DIR}src/static/css.tpl
rm -f ${TARGET_FILE}

CSS_FILES="${BASE_DIR}src/css/*"
FILEARY=()
for FILEPATH in ${CSS_FILES}; do
  if [ -f ${FILEPATH} ] ; then
    FILEARY+=("${FILEPATH}")
  fi
done

for i in ${FILEARY[@]}; do
  FILENAME=`basename ${i}`
  echo '{{define "'${FILENAME}'"}}' >> ${TARGET_FILE}
  cat ${i} >> ${TARGET_FILE}
  echo '{{end}}' >> ${TARGET_FILE}
done
