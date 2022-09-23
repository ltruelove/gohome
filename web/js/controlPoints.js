function ControlPointsViewModel(){
    var self = this;
    this.controlPoints = ko.observable([]);
    
    this.getControlPoints = function(){
        $.get('/controlPoint', function(response){
            self.controlPoints(response);
        });
    }

    this.deleteControlPoint = function(id) {
        console.log(id);
        if(confirm("Are you sure you want to delete this control point?")){
            $.ajax({
                url: '/controlPoint/' + id,
                type: 'DELETE',
                success: function(response){
                    console.log(response);
                    alert("Control point " + id + " deleted.");
                    self.getControlPoints();
                },
                data: null,
                contentType: "text/html"
            });
        };
    }

    $(window).on("load", function(){
        self.getControlPoints();
    });
}

ko.applyBindings(new ControlPointsViewModel());
