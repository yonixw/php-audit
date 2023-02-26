PHP Online
    https://onlinephp.io/
    
        https://github.com/docker-library/php/blob/master/8.2/bullseye/apache/Dockerfile
        
        https://github.com/docker-library/wordpress/blob/master/latest/php8.2/apache/Dockerfile
            php 8.0-apache => 8.0-apache-bullseye

    
        WP 5.8.1 version:
        https://github.com/docker-library/wordpress/blob/9954966feffdaf39082609816f896c2e3f75f0db/latest/php8.0/apache/Dockerfile

MYSQL:
    https://www.theairtips.com/post/mysql-101-audit-mysql-user-activities-with-mcafee-audit-plugin
Caddy:
    nginx log request:
        https://stackoverflow.com/a/7473603/1997873
    up to 100K ? https://github.com/caddyserver/caddy/issues/1015#issuecomment-239348304
        https://github.com/caddyserver/caddy/issues/1015#issuecomment-1035087612
    NodeJS parser and logger 
        Password mask
Outbound Connections:
    IP/DNS?
Files, glob, + http(s), ftp(s): 
    fread, fopen, file_get_contents, require, file(), curl_init, readfile, fgets, fgetss ...
            Common stuff? looks like all use php_stream_context
                how to log filename into CSV?
                    https://stackoverflow.com/questions/3531703/how-do-i-log-errors-and-warnings-into-a-file
                    libxml2 with --with-mem-debug
                        xmlMemoryDump
                
                PHP_FUNCTION(error_log)
                    https://www.educative.io/blog/concatenate-string-c
                    
                Stack info:
                    json_encode (debug_backtrace());
                    
                    zval backtrace;
                    zend_fetch_debug_backtrace(&backtrace, 1, 0, 0);
                    zend_string *str = zend_trace_to_string(Z_ARRVAL(backtrace), /* include_main */ false);
                        
            https://developer.ibm.com/articles/os-php-readfiles/
            php_stream *stream; -> 122 hits
            php_stream_open_wrapper_ex
    // print_r(stream_get_wrappers()); -> All builtins
    
    Modify stream_get_wrappers to return values?
        Block other from redefining? from registerning new?
    
    php_register_url_stream_wrapper
        php_register_url_stream_wrapper("compress.zlib", &php_stream_gzip_wrapper);
    php-8.2.3\ext\standard\streamsfuncs.c
    
    
    File watcher with hash, exclude wordpress_data - no ... use streams
    Or stream in php?
        https://stackoverflow.com/questions/5044925/php-hook-function
        
   static function wrap()
    {
        // self::PROTOCOL = "file";
        stream_wrapper_unregister(self::PROTOCOL);
        stream_wrapper_register(self::PROTOCOL, __CLASS__);
    }
    
    proxy class: https://gist.github.com/treffynnon/724314/63dd6ecc88b98860c684125aa86230e0944a34d8
    
    Wrappers:
        https://www.php.net/manual/en/class.streamwrapper.php

Save CSV:
    $profiling = fopen('profiling.csv', 'w');
    fputcsv($profiling, [getMethod(), $end - $begin]);


Custom build with replacement
    php-8.2.3.tar\mysqli\mysqli.stub.php
        see types, no connection  but msqli type...
        
    Add before all: https://stackoverflow.com/a/18075712/1997873

    in docker-php-source:
        todo: [^_\-a-zA-Z0-9]
        find ./ -type f -exec sed -i 's/xxxxxxxxxxxxmysqli_connect_errno/mysqli230223_connect_errno/g' {} \;
    
    in php/wordpress:
        docker build -t php:8.0-apache-230223 8.0/bullseye/apache/
        docker build -t wordpress:5.8.1-apache-230223 latest/php8.0/apache/Dockerfile
        
        docker tag php:8.0-apache-230223 yonixw/php-audit:8.0-apache-230223
        docker push yonixw/php-audit:8.0-apache-230223

    mysqli_connect_errno(): int {}
    
    
    check with new msqli()
        <?php
        $mysqli = new mysqli("localhost","my_user","my_password","my_db");

        // Check connection
        if ($mysqli -> connect_errno) {
          echo "Failed to connect to MySQL: " . $mysqli -> connect_error;
          exit();
        }
        ?>

    

    mysqli_init() -> $msqli with same wrapped stuff

    mysqli_debug($options) : true
    mysqli_dump_debug_info(connection): bool {}
    
    mysqli_real_connect
    mysqli_connect
    
    mysqli_execute_query(connection, query,?array $params = null)
    mysqli_real_query(connection, query)
    mysqli_multi_query(connection, query)
    mysqli_query(connection, query, resultmode)
    
    mysqli_stmt_execute
    mysqli_stmt_fetch
    mysqli_stmt_attr_set
    mysqli_stmt_prepare
    mysqli_stmt_reset
    
    mysqli_prepare(connection, query)
    mysqli_begin_transaction
    mysqli_savepoint
    mysqli_rollback
    mysqli_select_db
    
    mysqli_get_host_info
    mysqli_get_server_version
    mysqli_get_server_info
    mysqli_get_client_version
    mysqli_info
    mysqli_kill
    mysqli_options
    mysqli_real_escape_string
    mysqli_get_client_stats
    mysqli_get_connection_stats
    mysqli_set_charset
    mysqli_report


Eval catched by namespace trick?
    NOOOOO - uses global!
    
    All namespaces used:
    grep -rwshEo '^namespace.*?;' wordpress_data/ | awk -F"\\" '{print $1}' | awk -F" |;" '{print $2}' | sort | uniq

    Direct call used (that will be mocked):
        * my sql
        * eval
        grep -rwsnEo ' \\mysql_' wordpress_data/


    Mocks:
        https://akrabat.com/replacing-a-built-in-php-function-when-testing-a-component/
        https://github.com/php-mock/php-mock
        https://patchwork2.org/examples/