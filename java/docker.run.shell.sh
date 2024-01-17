#!/bin/bash
domain=$1
context=$2
configPath="/data/webapps/$domain/$context/conf"
ip=`env|grep MY_POD_IP|awk -F"=" '{print $2}'`
workDir="/data/webapps/$domain/$context"
menLimitBytes=`env|grep MY_MEM_LIMIT|awk -F"=" '{print $2}'`
menLimit=`echo "$menLimitBytes / 1024 / 1024 * 0.8" |bc|sed "s/\..*$//"`
namespaceName=`env|grep MY_POD_NAMESPACE|awk -F"=" '{print $2}'`
packageName=`env|grep MY_PACKAGE_NAME|awk -F"=" '{print $2}'`

mkdir -p /data/weblog/java/$domain
chmod -R 755 /data/weblog/java/$domain
touch /data/weblog/java/$domain/jb.log
chown -R www-data:www-data /data/weblog/java/$domain
mkdir -p /data/services/java_base/$domain

cd /data/webapps/${domain}
[ -d /data/webapps/${domain}/${context} ] || mkdir -p /data/webapps/${domain}/${context}

if [ -f "${context}.war" ];then
    file_name="${context}.war"
    mv $file_name /data/webapps/${domain}/${context}/
    cd /data/webapps/${domain}/${context}/
    unzip $file_name
    rm -f $file_name
    chown www-data.www-data /data/webapps/${domain}/${context} -R
fi

export MALLOC_ARENA_MAX=4
[ -f $workDir/META-INF/MANIFEST.MF ] && mainClass=`sed -r 's/\r//g' $workDir/META-INF/MANIFEST.MF | grep "Main-Class" | cut -d':' -f2 | sed 's/^[ \t]*//'`

if [ -f "${configPath}/dragon.javaOpts" ] && [ -s "${configPath}/dragon.javaOpts" ];then
    UserDesignOptions=`cat ${configPath}/dragon.javaOpts`
	xmsIsSetted=0
	xmxIsSetted=0
	echo "$UserDesignOptions" | grep "\-Xms" -q && xmsIsSetted=1
	echo "$UserDesignOptions" | grep "\-Xmx" -q && xmxIsSetted=1
	
	if [ "$xmsIsSetted" = "0" ] && [ "$xmxIsSetted" = "0" ];then
		DefaultJavaOpts="-Ddragon.bizName.projName=${namespaceName}____${packageName} -Ddragon.businessDomain=${packageName} -Ddragon.ip=${ip} -Xms${menLimit}m -Xmx${menLimit}m"
	elif [ "$xmsIsSetted" = "1" ] && [ "$xmxIsSetted" = "1" ];then
		DefaultJavaOpts="-Ddragon.bizName.projName=${namespaceName}____${packageName} -Ddragon.businessDomain=${packageName} -Ddragon.ip=${ip}"
	elif [ "$xmsIsSetted" = "0" ] && [ "$xmxIsSetted" = "1" ];then
		DefaultJavaOpts="-Ddragon.bizName.projName=${namespaceName}____${packageName} -Ddragon.businessDomain=${packageName} -Ddragon.ip=${ip} -Xms${menLimit}m -Xmx${menLimit}m"
	elif [ "$xmsIsSetted" = "1" ] && [ "$xmxIsSetted" = "0" ];then
		DefaultJavaOpts="-Ddragon.bizName.projName=${namespaceName}____${packageName} -Ddragon.businessDomain=${packageName} -Ddragon.ip=${ip} -Xms${menLimit}m -Xmx${menLimit}m"
	fi
    javaOpts="$DefaultJavaOpts $UserDesignOptions"
else
    DefaultJavaOpts="-Ddragon.bizName.projName=${namespaceName}____${packageName} -Ddragon.businessDomain=${packageName} -Ddragon.ip=${ip} -Xms${menLimit}m -Xmx${menLimit}m -Xmn512m -Xss256k -XX:+DisableExplicitGC -XX:+UseConcMarkSweepGC -XX:+CMSParallelRemarkEnabled -XX:+UseCMSCompactAtFullCollection -XX:LargePageSizeInBytes=128m -XX:+UseFastAccessorMethods -Djava.awt.headless=true -Djava.net.preferIPv4Stack=true -Duser.timezone=Asia/Shanghai -Dfile.encoding=UTF-8"

    javaOpts="$DefaultJavaOpts"
fi

[ -f ${configPath}/dragon.mainArgs ] && main_args=`cat ${configPath}/dragon.mainArgs`

echo "###########################CMD##############################"
echo "/usr/local/java/bin/java $javaOpts -classpath $workDir:$workDir/lib/* $mainClass $main_args"
echo "############################################################"

# Redirect stdout stderr
exec &> >(exec tee -a /data/weblog/java/$domain/jb.log)

# Start java app
exec /usr/local/java/bin/java $javaOpts -classpath $workDir:$workDir/lib/* $mainClass $main_args


