<!DOCTYPE html>
<html>
<head>
    <!-- 最新版jquery核心JavaScript文件 -->
    <script src="https://cdn.bootcss.com/jquery/3.2.1/jquery.js"></script>
    <!-- 最新版本的 Bootstrap 核心 CSS 文件 -->
    <link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css"
     integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
    <!-- 最新的 Bootstrap 核心 JavaScript 文件 -->
    <script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"
    integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
    <link rel="stylesheet" href="/static/css/bootstrap-datetimepicker.min.css" charset="UTF-8">
    <script src="/static/js/bootstrap-datetimepicker.min.js" charset="UTF-8"></script>

<script type="text/javascript">
$(document).ready(function(){
    $('.inputStartTime').datetimepicker({
        format:"yyyy-mm-dd",
        autoclose:true,
        todayBtn:true,
        minView:"month",
        pickerPosition:"bottom-left"
    });
    $('.inputEndTime').datetimepicker({
            format:"yyyy-mm-dd",
            autoclose:true,
            todayBtn:true,
            minView:"month",
            pickerPosition:"bottom-left"
        });

    $('#querySubmitButton').click(function(){
        var d = {};
        var t = $('#queryForm').serializeArray();
        $.each(t,function(){
            d[this.name] = this.value;
        });
        console.log("post json is: " + JSON.stringify(d));
        $.post("/query/scanner/result",JSON.stringify(d),function(result){
        console.log("get result is "+result);
            $("tr").remove(".result-data");
            var jsonResult = JSON.parse(result,null).scanner_results;
            for(var i in jsonResult){
                var one_tr = document.createElement("tr");
                var index = (parseInt(i)+1);
                one_tr.setAttribute('class','result-data');
                one_tr.innerHTML="<td>"+index+"</td>"+"<td>"+jsonResult[i].device_id+"</td>"
                +"<td>"+jsonResult[i].user_id+"</td>"
                +"<td>"+jsonResult[i].user_name+"</td>"
                +"<td>"+jsonResult[i].experiment_name+"</td>"
                +"<td x:str>"+jsonResult[i].total_algae_density+"</td>"
                +"<td>"+jsonResult[i].total_algae_count+"</td>"
                +"<td>"+jsonResult[i].advantage_algae_name+"</td>"
                +"<td x:str>"+jsonResult[i].advantage_algae_density+"</td>"
                +"<td>"+jsonResult[i].advantage_algae_percent+"</td>"
                +"<td>"+jsonResult[i].scanner_sample_volume+"</td>"
                +"<td>"+jsonResult[i].sample_volume+"</td>";
                $(".result-table").append(one_tr);
            }
        });
    });
});
</script>
</head>
<body>
   <form class="form-inline" id="queryForm">
    <div class="form-group">
        <label for="inputDeviceNo">设备编号</label>
        <input type="text" name="device_id" class="form-control" id="inputDeviceNo" placeholder="设备编号">
    </div>
    <div class="form-group">
            <label for="inputStartTime">起始时间</label>
            <div class="input-append date form_date form-group inputStartTime" data-link-field="inputStartTime" date-date-format="yyyy-mm-dd" date-picker-position="bottom-left">
                <input size="16" name="start_date" class="form-control" type="text" value="" required readonly>
                <!-- <span class="glyphicon glyphicon-calendar"></span> -->
                <span class="add-on"><i class="icon-th"></i></span>
            </div>
    </div>
    <div class="form-group">
            <label for="inputEndTime">{{.pageNums}}</label>
            <div class="input-append date form_date inputEndTime form-group" date-date-format="yyyy-mm-dd">
                <input size="16" name="end_date" class="form-control" type="text" value="" id="inputEndTime" required readonly>
                <span class="add-on"><i class="icon-th"></i></span>
            </div>
    </div>
    <button type="button" class="btn btn-primary" id="querySubmitButton">查询</button>
   </form>
   <br></br>
   <div class="table-responsive">
     <table class="table table-striped table-condensed result-table">
            <tr class="result-head">
                     <td>序号</td>
                     <td>设备编号</td>
                     <td>用户编号</td>
                     <td>用户昵称</td>
                     <td>实验名称</td>
                     <td>总藻密度</td>
                     <td>总藻数量</td>
                     <td>优势藻种</td>
                     <td>优势藻密度</td>
                     <td>优势藻占比</td>
                     <td>扫样容量(l)</td>
                     <td>取样容量(ml)</td>

           </tr>
        {{range $key,$val := .currentResult}}
            <tr class="result-data">
                <td>{{$key}}</td>
                <td>{{$val.DeviceId}}</td>
                <td>{{$val.UserId}}</td>
                <td>{{$val.UserName}}</td>
                <td>{{$val.ExperimentName}}</td>
                <td>{{$val.TotalAlgaeDensity}}</td>
                <td>{{$val.TotalAlgaeCount}}</td>
                <td>{{$val.AdvantageAlgaeName}}</td>
                <td>{{$val.AdvantageAlgaeDensity}}</td>
                <td>{{$val.AdvantageAlgaePercent}}</td>
                <td>{{$val.ScannerSampleVolume}}</td>
                <td>{{$val.SampleVolume}}</td>

            </tr>
        {{end}}
     </table>
   </div>
</body>

</html>