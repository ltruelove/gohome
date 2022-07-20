// Requires jQuery!!!

/*
This script is intended to be requested by the GoHome Control Point registration form
You MUST set the "apiHost" variable to the address of the GoHome API (with port)
in order for this to work properly.
*/

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

        registerControlPoint();
    });
});

function registerControlPoint() {
    var controlPoint = {
        "Name" : $('#controlPointName').val(),
        "IpAddress" : $('#controlPointIp').val(),
        "Mac" : $('#mac').val()
    }

    $.ajax({
        type: "POST",
        url: apiHost + "/controlPoint/register",
        data: JSON.stringify(controlPoint),
        success: registerSuccess,
        error: registerError,
        dataType: "json"
    });
}

async function registerSuccess(data){
    console.log(data);

    var res = await saveControlPointId(data.Id);
    console.log(res);
    await sleep(200);

    res = await callRestart();
    console.log(res);
}

function registerError(request, status, error){
    console.log(request);
    console.log(status);
    console.log(error);
}

function saveControlPointId(controlPointId){
    return $.ajax({
        type: "POST",
        url: "/setControlPointId",
        data: "controlPointId=" + controlPointId,
        success: genericSuccess,
        error: genericFailure,
        dataType: "text"
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