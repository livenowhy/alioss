
accessid = ''
accesskey = ''
host = ''
policyBase64 = ''
signature = ''
callbackbody = ''
filename = ''
key = ''
expire = 0
now = timestamp = Date.parse(new Date()) / 1000; 

function send_request()
{
    var xmlhttp = null;
    if (window.XMLHttpRequest)
    {
        xmlhttp=new XMLHttpRequest();
    }
    else if (window.ActiveXObject)
    {
        xmlhttp=new ActiveXObject("Microsoft.XMLHTTP");
    }
  
    if (xmlhttp!=null)
    {
        // var data = {"actionType": "ActionType-test", "actionResourceId": "actionResourceId-test"}

        // UserAvatars
	// MirrorIcon
        var data = JSON.stringify(
            {
                "actionType": "UserAvatars",
                "actionResourceId": "39828489-1bf6-334b-acdb-6a15bbd7c5a3",
            }
        )
        token = 'eyJ1aWQiOiAiYTU4Mzg1ZTAtN2ZkYy0zNDU4LWE3Y2YtN2QzNzY0YjlhYzg0IiwgInVzZXJfb3JhZyI6ICJsaXV6aGFuZ3BlaSIsICJ0b2tlbmlkIjogIjZlZTllNWFhYTJhNWM4NjdiNDExYzQyOSIsICJzYWx0IjogIjc5ODNhY2MzNThjMGU3Yzk2ZWQzOThlYyIsICJleHBpcmVzIjogMTQ4NDAyOTg0NC4zOTYwNjYsICJ1c2VyX29yZ2EiOiAibGl1emhhbmdwZWkiLCAicm9sZV91dWlkIjogIjIwMCIsICJvcmdhX3V1aWQiOiAiYTU4Mzg1ZTAtN2ZkYy0zNDU4LWE3Y2YtN2QzNzY0YjlhYzg0IiwgInVzZXJfdXVpZCI6ICJhNTgzODVlMC03ZmRjLTM0NTgtYTdjZi03ZDM3NjRiOWFjODQiLCAiZW1haWwiOiAibGl1emhhbmdwZWlAMTI2LmNvbSIsICJ1c2VyX25hbWUiOiAibGl1emhhbmdwZWkifYwvt239BEWEy3mlxCoeiCU='
        // phpUrl = 'http://123.56.9.18:8765/policy'
        // phpUrl = 'http://0.0.0.0:8765/policy'
        // phpUrl = 'http://192.168.1.5:8765/policy'
        phpUrl = 'https://img.boxlinker.com/policy'
        // phpUrl = 'http://www.livenowhy.com:8765/policy'
        xmlhttp.open( "POST", phpUrl, false );
        xmlhttp.setRequestHeader("token", token)
        // xmlhttp.send( null );
        xmlhttp.send( data );
        return xmlhttp.responseText
    }
    else
    {
        alert("Your browser does not support XMLHTTP.");
    }
};

function get_signature()
{
    //可以判断当前expire是否超过了当前时间,如果超过了当前时间,就重新取一下.3s 做为缓冲
    now = timestamp = Date.parse(new Date()) / 1000; 
    console.log('get_signature ...');
    console.log('expire:' + expire.toString());
    console.log('now:', + now.toString())
    if (expire < now + 3)
    {
        console.log('get new sign')
        body = send_request()
        var obj = eval ("(" + body + ")");
        host = obj['host']
        policyBase64 = obj['policy']
        accessid = obj['accessid']
        signature = obj['signature']
        expire = parseInt(obj['expire'])
        callbackbody = obj['callback']

        key = obj['dir']

        console.log(host)
        return true;
    }
    return false;
};

function set_upload_param(up)
{
    var ret = get_signature()
    if (ret == true)
    {
        new_multipart_params = {
            'key' : key + '${filename}',
            'policy': policyBase64,
            'OSSAccessKeyId': accessid, 
            'success_action_status' : '200', //让服务端返回200,不然，默认会返回204
            'callback' : callbackbody,
            'signature': signature,
        };

        up.setOption({
            'url': host,
            'multipart_params': new_multipart_params
        });

        console.log('reset uploader')
        console.log(new_multipart_params)
        console.log(callbackbody)
        //uploader.start();
    }
}

var uploader = new plupload.Uploader({
	runtimes : 'html5,flash,silverlight,html4',
	browse_button : 'selectfiles', 
	container: document.getElementById('container'),
	flash_swf_url : 'lib/plupload-2.1.2/js/Moxie.swf',
	silverlight_xap_url : 'lib/plupload-2.1.2/js/Moxie.xap',

    url : 'http://oss.aliyuncs.com',

	init: {
		PostInit: function() {
			document.getElementById('ossfile').innerHTML = '';
			document.getElementById('postfiles').onclick = function() {
            set_upload_param(uploader);
            uploader.start();
            return false;
			};
		},

		FilesAdded: function(up, files) {
			plupload.each(files, function(file) {
				document.getElementById('ossfile').innerHTML += '<div id="' + file.id + '">' + file.name + ' (' + plupload.formatSize(file.size) + ')<b></b>'
				+'<div class="progress"><div class="progress-bar" style="width: 0%"></div></div>'
				+'</div>';
			});
		},

		UploadProgress: function(up, file) {
			var d = document.getElementById(file.id);
			d.getElementsByTagName('b')[0].innerHTML = '<span>' + file.percent + "%</span>";
            
            var prog = d.getElementsByTagName('div')[0];
			var progBar = prog.getElementsByTagName('div')[0]
			progBar.style.width= 2*file.percent+'px';
			progBar.setAttribute('aria-valuenow', file.percent);
		},

		FileUploaded: function(up, file, info) {
            console.log('uploaded')
            console.log('callback')
            console.log(info.status)
            set_upload_param(up);
            if (info.status == 200)
            {
                document.getElementById(file.id).getElementsByTagName('b')[0].innerHTML = 'success';
            }
            else
            {
                document.getElementById(file.id).getElementsByTagName('b')[0].innerHTML = info.response;
            } 
		},

		Error: function(up, err) {
            set_upload_param(up);
			document.getElementById('console').appendChild(document.createTextNode("\nError xml:" + err.response));
		}
	}
});

uploader.init();
