function DoorViewModel(){
    var self = this;
    this.pinCode = ko.observable();

    this.doorStatus = function() {}

    this.pinValidate = function() {
        var code = self.pinCode();
        $.post('/clickGarageDoorButton', JSON.stringify({pinCode: code}), function(response){
            if(response.IsValid){
                // Check door status
            }
        }, 'json');
    }
}

ko.applyBindings(new DoorViewModel());
