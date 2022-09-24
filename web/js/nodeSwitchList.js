function NodeSwitchesViewModel(){
    var self = this;
    this.switches = ko.observable([]);
    let id = $('#NodeId').val();
    
    this.getNodeSwitches = function(){
        $.get('/node/switchesByNode/' + id, function(response){
            self.switches(response);
        });
    }

    $(window).on("load", function(){
        self.getNodeSwitches();
    });
}

ko.applyBindings(new NodeSwitchesViewModel());
