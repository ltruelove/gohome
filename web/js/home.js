function HomeViewModel(){
    var self = this;
    this.lightIp = ko.observable();
    this.lightStatus = ko.observable();

    this.turnLightOn = function() { $.get('http://' + self.lightIp() + '/on'); }
    this.turnLightOff = function() { $.get('http://' + self.lightIp() + '/off'); }
    this.lightStatus = function() { $.get('http://' + self.lightIp() + '/status', function(data){
        console.log(data); 
    }, 'json')}
}

var model = new HomeViewModel();

$.get('/lightIp', function(data){
    model.lightIp(data.ip);

    ko.applyBindings(model);
},'json');
