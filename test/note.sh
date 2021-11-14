#!/bin/bash
# Author:louisehong4168
# Blog:https://fenghong.tech
# Time:2020-12-06 11:17:39
# Name:test/note.sh
# Description:This is a production script.
echo "### Usage" >> Note.md
echo "
\`\`\`sh
# 下载并安装dmos, dmos是个golang的二进制工具，直接下载拷贝到bin目录即可, release页面也可下载
$ wget -c https://sealyun.oss-cn-beijing.aliyuncs.com/latest/dmos && \\
    chmod +x dmos && mv dmos /usr/bin
\`\`\`
" >> Note.md
echo "### [amd64 下载地址]" >> Note.md
echo "[oss 下载地址](https://${BUCKETNAME:-sealyun}.${OSSENDPOINT:-oss-cn-beijing.aliyuncs.com}/${VERSION}/dmos)" >> Note.md
echo "[latest 版本 oss下载地址](https://${BUCKETNAME:-sealyun}.${OSSENDPOINT:-oss-cn-beijing.aliyuncs.com}/latest/dmos)" >> Note.md
echo "### [arm64 下载地址]" >> Note.md
echo "[oss 下载地址](https://${BUCKETNAME:-sealyun}.${OSSENDPOINT:-oss-cn-beijing.aliyuncs.com}/${VERSION}/dmos-arm64)" >> Note.md
echo "[latest 版本 oss下载地址](https://${BUCKETNAME:-sealyun}.${OSSENDPOINT:-oss-cn-beijing.aliyuncs.com}/latest/dmos-arm64)" >> Note.md
echo "### Docker images" >> Note.md
echo "- \`${DOCKER_REPO:-fanux/dmos}:${VERSION}\`" >> Note.md
echo "- \`${DOCKER_REPO:-fanux/dmos}:latest\`" >> Note.md

echo "
### CHANGELOG
[https://github.com/fanux/dmos/blob/develop/CHANGELOG.md#${VERSION}](https://github.com/fanux/dmos/blob/develop/CHANGELOG.md#${VERSION})
" >> Note.md
cat Note.md
