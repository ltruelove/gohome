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

            $('#switchTypes').hide();
            $('.switchOptions').hide();
        }else if (this.value == 'switch') {
            $('#switchTypes').show();
            $('.switchOptions').show();

            $('#sensorTypes').hide();
        }
    });

    $('#addTypeButton').click(function(event){
        event.preventDefault();
        addType();
    })

    $('#switchTypes').hide();
    $('.switchOptions').hide();

    getSensorTypes();
    getSwitchTypes();
    clearTypeNameAndPin();
});

function registerNode() {
    var registrationDto = {
        "Node" : {
            "Name" : $('#nodeName').val(),
            "Mac" : $('#mac').val()
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

function registerSuccess(data){
    console.log(data);
}

function registerError(request, status, error){
    console.log(request);
    console.log(status);
    console.log(error);
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

    if(name === "" || pinNumber === ""){
        alert("Name and pin are required");
        return;
    }

    var sensor = {
        "Name" : name,
        "SensorTypeId" : parseInt(sensorVal),
        "Pin" : parseInt(pinNumber)
    };

    sensors.push(sensor);
    clearTypeNameAndPin();
    drawSensorList();
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
    sensors.splice(index, 1);
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
        "SensorTypeId" : parseInt(switchVal),
        "Pin" : parseInt(pinNumber),
        "MomentaryPressDuration" : $('#MomentaryPressDuration').val(),
        "IsClosedOn" : $('input[name="sType"]:checked').val()
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

function getSensorTypes() {
    $.get(apiHost + "/sensorType", function( data ) {
        console.log(data);
        if(data === null || data.length < 1){
            alert("No sensor types found");
            return;
        }

        for(var i = 0; i < data.length; i++){
            var sensorType = data[i];
            $('<option/>', {value: sensorType.Id }).text(sensorType.Name).appendTo('#sensorTypeOptions');
        }
    });
}

function getSwitchTypes() {
    $.get(apiHost + "/switchType", function( data ) {
        console.log(data);
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
