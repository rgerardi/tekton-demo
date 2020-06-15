#!/bin/sh
ctr=$(buildah from fedora:32)
wrkdir=/usr/local/bin
version=6
buildah copy $ctr hellogo $wrkdir
buildah run $ctr chmod a+x $wrkdir/hellogo
buildah run $ctr chgrp -R 0 $wrkdir 
buildah run $ctr chmod -R g=u $wrkdir
buildah config --workingdir $wrkdir $ctr
buildah config --cmd $wrkdir/hellogo $ctr
buildah config --entrypoint $wrkdir/hellogo $ctr
buildah commit $ctr hellogo
buildah tag hellogo hellogo:$version
buildah push hellogo:$version docker://quay.io/rgerardi/hellogo:$version

