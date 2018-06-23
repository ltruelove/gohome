function SoilViewModel(){
    var self = this;
    this.pinCode = ko.observable();
    this.soilLevel = ko.observable(0);
    this.waterStatus = 'Off';

    this.doorStatus = function() {}

    $.get('/soil', function(response){
        console.log(response);
        self.soilLevel(response);
    });

    this.turnWaterOn = function() {
        var code = self.pinCode();
        $.post('/waterOn', JSON.stringify({pinCode: code}), function(response){
            if(response.IsValid){
                // Check door status
            }
        }, 'json');
    }

    this.turnWaterOff = function(){
        var code = self.pinCode();
        $.post('/waterOff', JSON.stringify({pinCode: code}), function(response){
            if(response.IsValid){
                // Check door status
            }
        }, 'json');
    }
}

ko.applyBindings(new SoilViewModel());
