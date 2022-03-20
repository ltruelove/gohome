function DoorViewModel(){
    var self = this;
    this.pinCode = ko.observable();
    this.humidity = ko.observable();
    this.celcius = ko.observable();
    this.fahrenheit = ko.observable();
    this.doorClosed = ko.observable();

    this.doorStatus = function() {}

    this.pinValidate = function() {
        var code = self.pinCode();
        $.post('/clickGarageDoorButton', JSON.stringify({pinCode: code}), function(response){
            if(response.IsValid){
                // Check door status
            }
        }, 'json');
    }
    
    this.getStatus = function(){
        $.get('/temps/garage', function(response){
            if(response.ErrorMessage){
                console.log(response.ErrorMessage);
            }else{
                self.humidity(response.humidity);
                self.celcius(response.celcius);
                self.fahrenheit(response.fahrenheit);
                self.doorClosed(response.doorClosed);
            }
        }).fail(function(response, d){
            if(response.responseText){
                console.log(JSON.parse(response.responseText));
            }
        });
    }
    
    $(window).on("load", function(){
        self.getStatus();
        setInterval(function(){
            self.getStatus();
        }, 10000);
    });
}

ko.applyBindings(new DoorViewModel());
