function TempsViewModel(){
    var self = this;
    this.rooms = ko.observable([]);
    this.kidsHistory = [];
    this.masterHistory = [];
    this.garageHistory = [];
    this.mainHistory = [];
    
    this.getRoomStatus = function(){
        $.get('/temps/all', function(response){
            self.rooms(response);
        });
    }

    $(window).on("load", function(){
        self.getRoomStatus();

        setInterval(function(){
            self.getRoomStatus();
        }, 10000);
    });
}

ko.applyBindings(new TempsViewModel());
