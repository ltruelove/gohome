// Requires jQuery!!!

/*
This script is intended to be requested by the GoHome Node registration form
You MUST set the "apiHost" variable to the address of the GoHome API (with port)
in order for this to work properly.
*/
var sensorTypes = [];
var switchTypes = [];
var sensors = [];
var switches = [];

const DEF_DELAY = 1000;

function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms || DEF_DELAY));
}

$(document).ready(function() {
    if(typeof(apiHost) === 'undefined'){
        alert("Host is not set!");
        return;
    }
    console.log(apiHost);

    $("form").submit(function(e){
        e.preventDefault(e);

        registerNode();
    });

    $('input[name="sType"]').change(function() {
        if (this.value == 'sensor') {
            $('#sensorTypes').show();
            $('.dhtTypes').show();

            $('#switchTypes').hide();
            $('.switchOptions').hide();
        }else if (this.value == 'switch') {
            $('#switchTypes').show();
            $('.switchOptions').show();

            $('#sensorTypes').hide();
            $('.dhtTypes').hide();
        }
    });

    $('#sensorTypeOptions').on('change', function(){
        checkSensorTypes(this.value);
    });

    $('#addTypeButton').click(function(event){
        event.preventDefault();
        addType();
    })

    $('#switchTypes').hide();
    $('.switchOptions').hide();
    $('.dhtTypes').hide();

    getControlPoints();
    getSensorTypes();
    getSwitchTypes();
    clearTypeNameAndPin();
    checkSensorTypes($('#sensorTypeOptions :selected').val());
});

function registerNode() {
    var selectedControlPointId = $('#nodeControlPoint :selected').val();

    if(!selectedControlPointId){
        alert("You must select a control point");
        return false;
    }

    var registrationDto = {
        "Node" : {
            "Name" : $('#nodeName').val(),
            "Mac" : $('#mac').val()
        },
        "ControlPoint" : {
            "Id" : parseInt($('#nodeControlPoint :selected').val())
        },
        "Sensors" : sensors,
        "Switches" : switches
    }

    $.ajax({
        type: "POST",
        url: apiHost + "/node/register",
        data: JSON.stringify(registrationDto),
        success: registerSuccess,
        error: registerError,
        dataType: "json"
    });
}

async function registerSuccess(data){
    console.log(data);

    var res = await saveNodeId(data.Node.Id);
    console.log(res);
    await sleep(200);

    var res = await saveControlPointMac(data.ControlPoint.Mac);
    console.log(res);
    await sleep(200);

    for(var i = 0; i < data.Sensors.length; i++){
        var sensor = data.Sensors[i];

        switch(sensor.SensorTypeId){
            case 1:
                res = await saveDht(sensor.Pin, sensor.DHTType);
                await sleep(200);
            break;
            case 2:
                res = await saveMoisture(sensor.Pin);
                await sleep(200);
            break;
            default:
            break;
        }
    }

    for(var i = 0; i < data.Switches.length; i++){
        var switchData = data.Switches[i];

        switch(switchData.SwitchTypeId){
            case 1:
                res = await saveMomentary(switchData.Pin);
                await sleep(200);
            break;
            case 2:
                res = await saveToggle(switchData.Pin);
                await sleep(200);
            break;
            default:
            break;
        }
    }

    res = await callRestart();
    await sleep(200);
    console.log(res);
}

function saveControlPointMac(controlPointMac){
    return $.ajax({
        type: "POST",
        url: "/setControlPointMac",
        data: "controlPointMac=" + controlPointMac,
        success: genericSuccess,
        error: genericFailure,
        dataType: "json"
    });
}

function saveNodeId(nodeId){
    return $.ajax({
        type: "POST",
        url: "/setNodeId",
        data: "nodeId=" + nodeId,
        success: genericSuccess,
        error: genericFailure,
        dataType: "json"
    });
}

function saveDht(dhtPin, dhtType){
    return $.ajax({
        type: "POST",
        url: "/setDht",
        data: "pin=" + dhtPin + "&dhtType=" + dhtType,
        success: genericSuccess,
        error: genericFailure,
        dataType: "json"
    });
}

function saveMoisture(pin){
    return $.ajax({
        type: "POST",
        url: "/setMoisture",
        data: "pin=" + pin,
        success: genericSuccess,
        error: genericFailure,
        dataType: "json"
    });
}

function saveMomentary(pin){
    return $.ajax({
        type: "POST",
        url: "/setMomentary",
        data: "pin=" + pin,
        success: genericSuccess,
        error: genericFailure,
        dataType: "json"
    });
}

function saveToggle(pin){
    return $.ajax({
        type: "POST",
        url: "/setToggle",
        data: "pin=" + pin,
        success: genericSuccess,
        error: genericFailure,
        dataType: "json"
    });
}

function callRestart(){
    return $.get("/restart", function(data){
        console.log(data);
    });
}

function genericSuccess(data){
    console.log(data);
}

function genericFailure(req, msg, err){
    console.log(req);
    console.log(msg);
    console.log(err);
}

function registerError(request, status, error){
    console.log(request);
    console.log(status);
    console.log(error);
}

function checkSensorTypes(sensorVal){
    if($('#sensorTypeOptions').is(':hidden')){
        return;
    }

    if(sensorVal == 1){
        $('.dhtTypes').show();
    }else{
        $('.dhtTypes').hide();
    }
}

function clearTypeNameAndPin(){
    $('#typeName').val('');
    $('#pin').val('');
}

function addType(){
    var typeSelected = $('input[name="sType"]:checked');

    if(typeSelected.val() === 'sensor'){
        addSensor();
    }else{
        addSwitch();
    }
}

function addSensor(){
    var name = $('#typeName').val();
    var sensorVal = $('#sensorTypeOptions :selected').val();
    var pinNumber = $('#pin').val();
    var dhtType = $('input[name="dhtType"]:checked').val();

    if(name === "" || pinNumber === ""){
        alert("Name and pin are required");
        return;
    }

    var sensor = {
        "Name" : name,
        "SensorTypeId" : parseInt(sensorVal),
        "Pin" : parseInt(pinNumber),
        "DHTType" : parseInt(dhtType)
    };

    sensors.push(sensor);
    clearTypeNameAndPin();
    drawSensorList();
    $('#sensorTypeOptions option[value="'+ sensorVal +'"]').prop("disabled", true);
}

function drawSensorList(){
    const listDiv = $('#sensorList');
    listDiv.html('');

    for(var i = 0; i < sensors.length; i++){
        var sensor = sensors[i];
        var item = $("<div/>");
        item.html("<p>" + sensor.Name + ": Pin " + sensor.Pin + " &nbsp; <button onclick='removeSensor(" + i +")'>X</button></p>");
        item.appendTo(listDiv);
    }
}

function removeSensor(index){
    var removedSensor = sensors.splice(index, 1);

    $('#sensorTypeOptions option[value="'+ removedSensor.SensorTypeId +'"]').prop("disabled", false);

    drawSensorList();
}

function addSwitch(){
    var name = $('#typeName').val();
    var switchVal = $('#switchTypeOptions :selected').val();
    var pinNumber = $('#pin').val();

    if(name === "" || pinNumber === ""){
        alert("Name and pin are required");
        return;
    }

    var switchType = {
        "Name" : name,
        "SwitchTypeId" : parseInt(switchVal),
        "Pin" : parseInt(pinNumber),
        "MomentaryPressDuration" : parseInt($('#MomentaryPressDuration').val()),
        "IsClosedOn" : $('input[name="IsClosedOn"]:checked').val() !== "1" ? false : true
    };

    switches.push(switchType);
    clearTypeNameAndPin();
    drawSwitchList();
}

function drawSwitchList(){
    const listDiv = $('#switchList');
    listDiv.html('');

    for(var i = 0; i < switches.length; i++){
        var switchType = switches[i];
        var item = $("<div/>");
        item.html("<p>" + switchType.Name + ": Pin " + switchType.Pin + " &nbsp; <button onclick='removeSwitch(" + i +")'>X</button></p>");
        item.appendTo(listDiv);
    }
}

function removeSwitch(index){
    switches.splice(index, 1);
    drawSwitchList();
}

function getControlPoints() {
    $.get(apiHost + "/controlPoint/Available", function( data ) {
        if(data === null || data.length < 1){
            alert("No control points found. Please add a control point before adding a node.");
            return;
        }

        for(var i = 0; i < data.length; i++){
            var controlPoint = data[i];
            $('<option/>', {value: controlPoint.Id }).text(controlPoint.Name).appendTo('#nodeControlPoint');
        }
    });
}

function getSensorTypes() {
    $.get(apiHost + "/sensorType", function( data ) {
        if(data === null || data.length < 1){
            alert("No sensor types found");
            return;
        }

        for(var i = 0; i < data.length; i++){
            var sensorType = data[i];

            // if the device already has a pin defined for DHT, skip DHT
            if(dhtPin && dhtPin > 0){
                if(sensorType.Id === 1){
                    continue;
                }
            }
            $('<option/>', {value: sensorType.Id }).text(sensorType.Name).appendTo('#sensorTypeOptions');
        }

        checkSensorTypes($('#sensorTypeOptions :selected').val());
    });
}

function getSwitchTypes() {
    $.get(apiHost + "/switchType", function( data ) {
        if(data === null || data.length < 1){
            alert("No switch types found");
            return;
        }

        for(var i = 0; i < data.length; i++){
            var switchType = data[i];
            $('<option/>', {value: switchType.Id }).text(switchType.Name).appendTo('#switchTypeOptions');
        }
    });
}
