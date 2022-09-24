function NodeSwitchesViewModel(){
    var self = this;
    this.switches = ko.observable([]);
    let id = $('#NodeId').val();
    
    this.getNodeSwitches = function(){
        $.get('/node/switchesByNode/' + id, function(response){
            console.log(response);
            self.switches(response);
        });
    }

    this.toggleSwitch = function(toggleId){
        $.get('/node/switch/toggle/' + toggleId, function(response){
            console.log(response);
        });
    }

    $(window).on("load", function(){
        self.getNodeSwitches();
    });
}

ko.applyBindings(new NodeSwitchesViewModel());
