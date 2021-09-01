function TempsViewModel(){
    var self = this;
    this.kids = {"humidity":ko.observable(), "celcius":ko.observable(), "fahrenheit": ko.observable() };
    this.master = {"humidity":ko.observable(), "celcius":ko.observable(), "fahrenheit": ko.observable() };
    this.garage = {"humidity":ko.observable(), "celcius":ko.observable(), "fahrenheit": ko.observable() };
    this.main = {"humidity":ko.observable(), "celcius":ko.observable(), "fahrenheit": ko.observable() };
    
    this.getGarageStatus = function(){
        $.get('/garage', function(response){
            self.garage.humidity(response.humidity);
            self.garage.celcius(response.celcius);
            self.garage.fahrenheit(response.fahrenheit);
        });
    }
    
    this.getKidsStatus = function(){
        $.get('kids', function(response){
            self.kids.humidity(response.humidity);
            self.kids.celcius(response.celcius);
            self.kids.fahrenheit(response.fahrenheit);
        });
    }
    
    this.getMasterStatus = function(){
        $.get('/master', function(response){
            self.master.humidity(response.humidity);
            self.master.celcius(response.celcius);
            self.master.fahrenheit(response.fahrenheit);
        });
    }
    
    this.getMainStatus = function(){
        $.get('/main', function(response){
            self.main.humidity(response.humidity);
            self.main.celcius(response.celcius);
            self.main.fahrenheit(response.fahrenheit);
        });
    }

    this.fetchResults = function(){
        self.getGarageStatus();
        self.getKidsStatus();
        self.getMasterStatus();
        self.getMainStatus();
    }

    $(window).on("load", function(){
        self.fetchResults();

        setInterval(function(){
            self.fetchResults();
        }, 10000);
    });
}

ko.applyBindings(new TempsViewModel());
