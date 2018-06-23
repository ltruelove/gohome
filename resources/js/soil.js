function SoilViewModel(){
    var self = this;
    this.pinCode = ko.observable();
    this.soilLevel = ko.observable(0);
    this.waterStatus = ko.observable('unknown');

    this.doorStatus = function() {}

    $.get('/soil', function(response){
        self.soilLevel(response);
    });

    self.getStatus = function(){
        $.get('/waterStatus', function(response){
            self.waterStatus(response);
        });
    }

    self.getStatus();

    this.turnWaterOn = function() {
        var code = self.pinCode();
        $.post('/waterOn', JSON.stringify({pinCode: code}), function(response){
            console.log(response);
            self.getStatus();
        }, 'json');
    }

    this.turnWaterOff = function(){
        var code = self.pinCode();
        $.post('/waterOff', JSON.stringify({pinCode: code}), function(response){
            console.log(response);
            self.getStatus();
        }, 'json');
    }
}

ko.applyBindings(new SoilViewModel());
