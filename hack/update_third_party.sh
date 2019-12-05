#! /bin/bash

set -e

PROJECT_DIR=$(cd $(dirname $0)/..; pwd)
VERSION=$(grep aws-sdk-go-v2 ${PROJECT_DIR}/go.mod | cut -d " " -f 2)
TEMP_HEADER_TXT="/tmp/nifcloud_sdk_go_v2_header.txt"
TEMP_OUTPUT="/tmp/nifcloud_sdl_go_v2_output.go"

fork_from_aws_sdk_go_v2() {
    DOWNLOAD_URL=$1
    PAGE_URL=$2
    OUTPUT=$3
    curl -s ${DOWNLOAD_URL} -o ${OUTPUT}
    echo -e "// This code was forked from github.com/aws/aws-sdk-go-v2. DO NOT EDIT.\n// URL: ${PAGE_URL}\n" > ${TEMP_HEADER_TXT}
    cat ${TEMP_HEADER_TXT} ${OUTPUT} > ${TEMP_OUTPUT}
    mv ${TEMP_OUTPUT} ${OUTPUT}
    rm ${TEMP_HEADER_TXT}
}

update_third_party::aws_sdk_go_v2::internal::awsutil() {
    DOWNLOAD_URL="https://raw.githubusercontent.com/aws/aws-sdk-go-v2/${VERSION}/internal/awsutil/prettify.go"
    PAGE_URL="https://github.com/aws/aws-sdk-go-v2/tree/${VERSION}/internal/awsutil/prettify.go"
    OUTPUT="${PROJECT_DIR}/internal/nifcloudutil/prettify.go"
    fork_from_aws_sdk_go_v2 ${DOWNLOAD_URL} ${PAGE_URL} ${OUTPUT}
    sed -i "" -e "s/package awsutil/package nifcloudutil/g" ${OUTPUT}
}

update_third_party::aws_sdk_go_v2::aws::endpoints() {
    BASE_DOWNLOAD_URL="https://raw.githubusercontent.com/aws/aws-sdk-go-v2/${VERSION}/aws/endpoints/"
    BASE_PAGE_URL="https://github.com/aws/aws-sdk-go-v2/tree/${VERSION}/aws/endpoints/"
    BASE_OUTPUT="${PROJECT_DIR}/nifcloud/endpoints/"
    declare -a TARGETS=("endpoints.go" "v3model.go")

    for TARGET in ${TARGETS[@]}; do
        fork_from_aws_sdk_go_v2 ${BASE_DOWNLOAD_URL}${TARGET} ${BASE_PAGE_URL}${TARGET} ${BASE_OUTPUT}${TARGET}
    done
}

update_third_party::aws_sdk_go_v2::aws::endpoints
update_third_party::aws_sdk_go_v2::internal::awsutil
