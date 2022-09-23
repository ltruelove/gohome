function NodesViewModel(){
    var self = this;
    this.nodes = ko.observable([]);
    
    this.getNodes = function(){
        $.get('/node', function(response){
            self.nodes(response);
        });
    }

    this.deleteNode = function(id) {
        console.log(id);
        if(confirm("Are you sure you want to delete this node?")){
            $.ajax({
                url: '/node/' + id,
                type: 'DELETE',
                success: function(response){
                    console.log(response);
                    alert("Node " + id + " deleted.");
                    self.getNodes();
                },
                data: null,
                contentType: "text/html"
            });
        };
    }

    $(window).on("load", function(){
        self.getNodes();
    });
}

ko.applyBindings(new NodesViewModel());
