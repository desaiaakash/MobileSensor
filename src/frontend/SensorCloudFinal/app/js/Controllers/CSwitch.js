/**
 * Created by Pulkit on 03-05-2016.
 */

angular.module('CSwitch',['ngRoute','ngAnimate','ngCookies'])
.config(['$routeProvider', '$locationProvider','$httpProvider',
    function($routeProvider, $locationProvider,$httpProvider) {
            $httpProvider.defaults.headers.common = {};
            $httpProvider.defaults.headers.post = {};
            $httpProvider.defaults.headers.put = {};
            $httpProvider.defaults.headers.patch = {};

        $routeProvider
            .when('/Dashboard', {
                templateUrl: 'Dashboard.html',
                controller: "DashCtrl",
               // controllerAs: 'dash'
            })
            .when('/Profile', {
                templateUrl: 'profile.html',
                controller: "ProfileCtrl",
                //controllerAs: 'prof'
            })
            .when('/SensorManagement', {
                templateUrl: 'Sensor Management.html',
                controller: 'SensorCtrl'

            })
            .when('/Reports', {
                templateUrl: 'Reports.html',
                controller:'ReportCtrl'

            })
            .when('/Maps', {
                templateUrl: 'Maps.html',
                controller:'MapCtrl'

            })
            .when('/Billing', {
                templateUrl: 'Billing.html',
                controller:'BillingCtrl'

            })
            .when('/Subscription',{
                templateUrl: 'Subscriptions.html',
                controller:'SubscriptionsCtrl'
            })
            .otherwise({
            redirectTo:'Dashboard.html'
        });

       // $locationProvider.html5Mode(true);
    }])
//Mounika
    .controller('ReportCtrl',['$scope', function($scope){
        $scope.message ="Successful Report";

        /*  var request = $http({
         method: "get",
         url: '/user/billing/'+$scope.user.uName+'/'+$scope.user.pass,
         });

         request.success(
         function(response) {
         $scope.list = response ;

         }); */

        $scope.list=[{
            sensorid: "2114e2dqwdqw1",
            location: "san jose",
            type: "Temperature",
            typecode: 1,
            value : 13
        }, {
            sensorid: "2114e2dqwdqw2",
            location: "san jose",
            type: "Pressure",
            typecode: 2,
            value : 39
        },
             {
                sensorid: "2114e2dqwdqw3",
                location: "san jose",
                type: "Humidity",
                typecode: 3,
                value : 82
            }
            , {
                sensorid: "2114e2dqwdqw4",
                location: "san jose",
                type: "Air Quality Index",
                typecode: 4,
                value : 44
            }
            , {
                sensorid: "2114e2dqwdqg1",
                location: "Fresno",
                type: "Pressure",
                typecode: 2,
                value : 28
            },
            { sensorid: "2114e2dqwdqg2",
                location: "Fresno",
                type: "Temperature",
                typecode: 1,
                value : 13
            },
            { sensorid: "2114e2dqwdqg3",
                location: "Fresno",
                type: "Humidity",
                typecode: 3,
                value : 77
            },
            {sensorid: "2114e2dqwdqg4",
                location: "Fresno",
                type: "Air Quality Index",
                typecode: 4,
                value : 43
            }
            , {
                sensorid: "2114e2dqwdqf1",
                location: "Apple Valley",
                type: "Pressure",
                typecode: 2,
                value : 20
            },
            {
                sensorid: "2114e2dqwdqf2",
                location: "Apple Valley",
                type: "Temperature",
                typecode: 1,
                value : 20
            },
            {
                sensorid: "2114e2dqwdqf3",
                location: "Apple Valley",
                type: "Humidity",
                typecode: 3,
                value : 80
            },
            {
                sensorid: "2114e2dqwdqf4",
                location: "Apple Valley",
                type: "Air Quality Index",
                typecode: 4,
                value : 39
            },{
                sensorid: "2114e2dqwdqk4",
                location: "Goleta",
                type: "Pressure",
                typecode: 2,
                value : 39
            },
            {
                sensorid: "2114e2dqwdqk3",
                location: "Goleta",
                type: "Temperature",
                typecode: 1,
                value : 40
            },
            {
                sensorid: "2114e2dqwdqk2",
                location: "Goleta",
                type: "Humidity",
                typecode: 3,
                value : 72
            },
            {
                sensorid: "2114e2dqwdqk1",
                location: "Goleta",
                type: "Air Quality Index",
                typecode: 4,
                value : 43
            }];

    }])
    .controller('BillingCtrl',['$scope', function($scope){
        $scope.message ="Successful Billing Details";

        /*  var request = $http({
         method: "get",
         url: '/user/billing/'+$scope.user.uName+'/'+$scope.user.pass,
         });

         request.success(
         function(response) {
         $scope.list = response ;
         // $scope.list.costs = response.cost*response.hoursConsumed ;

         }); */

        $scope.list=[{
            sensorid: "2114e2dqwdqw1",
            location: "san jose",
            type: "Temperature",
            typecode: 1,
            cost: 0.2,
            hoursConsumed: 6,
            costs : 1.2
        }, {
            sensorid: "2114e2dqwdqw2",
            location: "san jose",
            type: "Pressure",
            typecode: 2,
            cost: 0.2,
            hoursConsumed: 6,
            costs : 1.2
        },
            {
                sensorid: "2114e2dqwdqw3",
                location: "san jose",
                type: "Humidity",
                typecode: 3,
                cost: 0.2,
                hoursConsumed: 6,
                costs : 1.2
            },
            {
                sensorid: "2114e2dqwdqw4",
                location: "san jose",
                type: "Air Quality Index",
                typecode: 4,
                cost: 0.2,
                hoursConsumed: 6,
                costs : 1.2
            },
            {
                sensorid: "2114e2dqwdqg1",
                location: "Fresno",
                type: "Pressure",
                typecode: 2,
                cost: 0.2,
                hoursConsumed: 6,
                costs : 1.2
            },
            { sensorid: "2114e2dqwdqg2",
                location: "Fresno",
                type: "Temperature",
                typecode: 1,
                cost: 0.2,
                hoursConsumed: 6,
                costs : 1.2
            },
            { sensorid: "2114e2dqwdqg3",
                location: "Fresno",
                type: "Humidity",
                typecode: 3,
                cost: 0.2,
                hoursConsumed: 6,
                costs : 1.2
            },
            {sensorid: "2114e2dqwdqg4",
                location: "Fresno",
                type: "Air Quality Index",
                typecode: 4,
                cost: 0.2,
                hoursConsumed: 6,
                costs : 1.2
            },
            {
                sensorid: "2114e2dqwdqf1",
                location: "Apple Valley",
                type: "Pressure",
                typecode: 2,
                cost: 0.2,
                hoursConsumed: 6,
                costs : 1.2
            },
            {
                sensorid: "2114e2dqwdqf2",
                location: "Apple Valley",
                type: "Temperature",
                typecode: 1,
                cost: 0.2,
                hoursConsumed: 6,
                costs : 1.2
            },
            {
                sensorid: "2114e2dqwdqf3",
                location: "Apple Valley",
                type: "Humidity",
                typecode: 3,
                cost: 0.2,
                hoursConsumed: 6,
                costs : 1.2
            },
            {
                sensorid: "2114e2dqwdqf4",
                location: "Apple Valley",
                type: "Air Quality Index",
                typecode: 4,
                cost: 0.2,
                hoursConsumed: 6,
                costs : 1.2
            },{
                sensorid: "2114e2dqwdqk4",
                location: "Goleta",
                type: "Pressure",
                typecode: 2,
                cost: 0.2,
                hoursConsumed: 6,
                costs : 1.2
            },
            {
                sensorid: "2114e2dqwdqk3",
                location: "Goleta",
                type: "Temperature",
                typecode: 1,
                cost: 0.2,
                hoursConsumed: 6,
                costs : 1.2
            },
            {
                sensorid: "2114e2dqwdqk2",
                location: "Goleta",
                type: "Humidity",
                typecode: 3,
                cost: 0.2,
                hoursConsumed: 6,
                costs : 1.2
            },
            {
                sensorid: "2114e2dqwdqk1",
                location: "Goleta",
                type: "Air Quality Index",
                typecode: 4,
                cost: 0.2,
                hoursConsumed: 6,
                costs : 1.2
            }];
        $scope.total = function(){
            var total = 0;
            for(count=0;count<$scope.list.length;count++){
                total += $scope.list[count].costs + $scope.list[count].costs;
            }
            return total;
        }

    }])
    .controller('DashCtrl',['$scope', function( $scope){
        $scope.message ="Successful Dash";

        Highcharts.chart('containerss', {
            title: {
                text: 'Temperature Data'
            },

            xAxis: {
                categories: ['1', '2', '3', '4', '5']
            },

            series: [{
                data: [13, 15, 25, 10],

            }]
        });


        Highcharts.chart('containers', {
            title: {
                text: 'Temperature Data'
            },
            chart: {
                type: 'column',
                margin: 75,
                options3d: {
                    enabled: true,
                    alpha: 15,
                    beta: 15,
                    depth: 50,
                    viewDistance: 25
                }
            },
            xAxis: {
                categories: ['San Jose', 'Fresno', 'Goleta', 'Apple Valley']
            },

            plotOptions: {
                column: {
                    depth: 25
                },
                series:{
                    color: '#0080ff'
                }
            },
            series: [{
                name:"Temprature",
                data: [13, 10, 26, 32],

            }]
        });

        Highcharts.chart('pressure', {
            title: {
                text: 'Pressure Data'
            },
            chart: {
                type: 'column',
                margin: 75,
                options3d: {
                    enabled: true,
                    alpha: 15,
                    beta: 15,
                    depth: 50,
                    viewDistance: 25
                }
            },xAxis: {
                categories: ['San Jose', 'Fresno', 'Goleta', 'Apple Valley']
            },

            plotOptions: {
                column: {
                    depth: 25
                },
                series:{
                    color: '#ff4000'
                }
            },
            series: [{
                name:"Pressure",
                data: [30, 20, 29, 34],

            }]
        });

        Highcharts.chart('humidity', {
            title: {
                text: 'Humidity Data'
            },
            chart: {
                type: 'column',
                margin: 75,
                options3d: {
                    enabled: true,
                    alpha: 15,
                    beta: 15,
                    depth: 50,
                    viewDistance: 25
                }
            },
            plotOptions: {
                column: {
                    depth: 25
                },
                series:{
                    color: '#4000ff'
                }
            },xAxis: {
                categories: ['San Jose', 'Fresno', 'Goleta', 'Apple Valley']
            },

            series: [{
                name:"Humidity",
                data: [71, 60, 80, 73],

            }]
        });

        Highcharts.chart('airqualityindex', {
            title: {
                text: 'Air Quality Index'
            },
            chart: {
                type: 'column',
                margin: 75,
                options3d: {
                    enabled: true,
                    alpha: 15,
                    beta: 15,
                    depth: 50,
                    viewDistance: 25
                }
            },xAxis: {
                categories: ['San Jose', 'Fresno', 'Goleta', 'Apple Valley']
            },

            plotOptions: {
                column: {
                    depth: 25
                },
                series:{
                    color: '#39e600'
                }
            },
            series: [{
                name:"Air Quality",
                data: [46, 40, 36, 42],

            }]
        });


    }])
    .controller('ProfileCtrl',['$scope', function($scope){
        $scope.message ="Successful Profile";
    }])
    .controller('MapCtrl',['$scope','$http',function($scope,$http){
        var cities = [
            {
                city : 'San Jose',
                desc : '',
                lat : 37.,
                long : 121.867905
            },
            {
                city : 'Fresno',
                desc : 'This city is aiiiiite!',
                lat : 36.7468,
                long : -119.7726
            },
            {
                city : 'Apple Valley',
                desc : 'This is the second best city in the world!',
                lat : 34.397844946449865,
                long : -117.15442657470703
            },
            {
                city : 'Goleta',
                desc : 'This city is live!',
                lat : 34.72355492704219,
                long : -120.20862579345703
            }

        ];
        function showMaps(){
            var mapOptions = {
                center: new google.maps.LatLng(37.338208, -121.886329),
                zoom: 6,
                mapTypeId: google.maps.MapTypeId.ROADMAP

                };
            new google.maps.Map(document.getElementById("dvMap"), mapOptions);

            $scope.map = new google.maps.Map(document.getElementById("dvMap"), mapOptions);

            //Get data through a fn and put in here using markers

            $scope.markers = [];

            var infoWindow = new google.maps.InfoWindow();

            var createMarker = function (info){

                var marker = new google.maps.Marker({
                    map: $scope.map,
                    position: new google.maps.LatLng(info.lat, info.long),
                    title: info.city
                });
                marker.content = '<div class="infoWindowContent">' + info.desc + '</div>';

                google.maps.event.addListener(marker, 'click', function(){
                    infoWindow.setContent('<h2>' + marker.title + '</h2>' + marker.content);
                    infoWindow.open($scope.map, marker);
                });

                $scope.markers.push(marker);

            }

            for (i = 0; i < cities.length; i++){
                createMarker(cities[i]);
            }


        }
        showMaps();

    }])
    .controller('SensorCtrl',['$scope','$http','$window','$cookies','$timeout',function($scope,$http,$window,$cookies,$timeout){
        $scope.alShow=false;
        var userN =$cookies.get('uName');
        var userP =$cookies.get('uPass');
        $scope.alHide= true;
        function successFunction() {

                var mapOptions = {
                    center: new google.maps.LatLng(37.338208, -121.886329),
                    zoom: 6,
                    mapTypeId: google.maps.MapTypeId.ROADMAP
                };
              var map = new google.maps.Map(document.getElementById("dvMap1"), mapOptions);
                marker = new google.maps.Marker({
                     map: map,
                     draggable: true,
                     animation: google.maps.Animation.DROP,
                     position: {lat: 37.338208, lng: -121.886329}
            });
                    google.maps.event.addListener(marker, "dragend", function (e) {
                    $scope.latitude = e.latLng.lat();
                    $scope.longitude = e.latLng.lng();


                });

        }
        successFunction();
        $scope.sTypes = {
            1: "Temperature",
            2: "Pressure",
            3: "Humidity",
            4: "Air Quality Index"
        };
        $scope.getCity=function() {
            $http({
                method: 'GET',
                url: "https://maps.googleapis.com/maps/api/geocode/json?latlng=" + $scope.latitude + "," + $scope.longitude

            }).then(function (resp) {
                    if (resp.data.status != "ZERO_RESULTS") {
                        console.log(resp.data);
                        var str = resp.data['results'][1]['formatted_address'];
                        var lastIndex = str.lastIndexOf(",");
                        str = str.substring(0, lastIndex);
                        var pieces = str.split(/[\,]+/);
                        var newcity = pieces[pieces.length - 2];
                        $scope.city = newcity;
                    }
                    else {
                        $window.alert("Enter a valid location");
                    }
                },
                function (resp) {
                    this.errMsg = resp.data;
                });
        }
//Rest not working here backend doesn't support delete
        $scope.removeSensor=function(){
            console.log('http://127.0.0.1:3503/sensor/remove/PD/123/'+ $scope.sensorId);
            $http({
                method:'DELETE',
                url: 'http://127.0.0.1:3503/sensor/remove/PD/123/'+ $scope.sensorId , //change PD to userN and 123 to userP
                headers:undefined,
            }).success(function(resp){
                //was it successful
                console.log(resp);
                if(resp.data.Result=='true'){
                    console.log('Sensor Removed Successfully');
                }
            })
                .error(function(resp){
                console.log("COuld nt remove sensor"+resp);
            })
        }
//For hiding alerts on Sensor Management page
        $scope.hideAlert = function(){
            $scope.alShow=false;
        }
       /* $scope.addSensor=function() {
            //Temp Code all these will go in the implementation and assigned dynamically
            $scope.alType = "success";
            $scope.alShow = true;
            $scope.alMessage ="Sensor Add was successful."


        }*/
        //Final Code
        $scope.addSensor=function(){
            $http({
                method: 'POST',
                url:'http://127.0.0.1:3503/sensor/add/',
                headers:undefined,
                data:{

                    "Username": "PD" ,      //Change PD to $cookies.get('uName')
                    "Password": '123' ,  //$cookies.get('uPass')
                    "Location": $scope.city,
                    "Type": $scope.sTypes[$scope.sType] ,
                    "Typecode": $scope.sType,
                    "Coordinate": {
                        "Lat": $scope.latitude,
                        "Lng": $scope.longitude
                    }

                }
            }).then(function(resp){
                console.log(resp.data);
                if (resp.data.typecode == $scope.sType) {
                    $scope.alShow =true;
                    $scope.alType = "success";
                    $scope.alMessage ="Sensor Add was successful.";
                    //$scope.alert = "Your Sensor was added successfully";
                    console.log('Sensor ADDED Successfully');
                }
            },function(resp){
                $scope.alShow=true;
                $scope.alType = "danger";
                $scope.alMessage ="Sensor Add was Unsuccessful.";
                console.log("COuld nt add sensor "+resp.data);
            })
        }
        $scope.subscribe = function(){
            $http({
                method: 'GET',
                url:'http://127.0.0.1:3503/sensor/status/PD/123/'+$scope.sensorId2+'/1',
                headers:undefined,
            }).then(function(resp){
            //Show alert that successfully Subscribe
                console.log("Subscription Successful");
            }, function(resp){

                console.log("Subscription Failed"+resp);
            }
            )

        }
        $scope.unsubscribe = function(){
            $http({
                method: 'GET',
                url:'http://127.0.0.1:3503/sensor/status/PD/123/'+$scope.sensorId3+'/0',
                headers:undefined,
            }).then(function(resp){
                    //Show alert that successfully Subscribe
                    console.log("UnSubscription Successful");
                }, function(resp){
                    console.log("UnSubscription Failed"+resp);
                }
            )

        }

    }])
.controller('SubscriptionsCtrl',['$scope', function($scope){
    $scope.message ="Successful Sub";

    /*  var request = $http({
     method: "get",
     url: '/user/billing/'+$scope.user.uName+'/'+$scope.user.pass,
     });

     request.success(
     function(response) {
     $scope.list = response ;

     }); */


    $scope.list=[{
        sensorid: "2114e2dqwdqw1",
        location: "san jose",
        type: "Temperature",
        typecode: 1,
        cost: 0.2,
        hoursConsumed: 6,
        created_date : "05/10/2016",
        status : "Active",
        end_date : "-"
    }, {
        sensorid: "2114e2dqwdqw2",
        location: "san jose",
        type: "Pressure",
        typecode: 2,
        cost: 0.2,
        hoursConsumed: 6,
        created_date : "05/10/2016",
        status : "Active",
        end_date : "-"
    },
        {
            sensorid: "2114e2dqwdqw3",
            location: "san jose",
            type: "Humidity",
            typecode: 3,
            cost: 0.2,
            hoursConsumed: 6,
            created_date : "05/10/2016",
            status : "Active",
            end_date : "-"
        },
        {
            sensorid: "2114e2dqwdqw4",
            location: "san jose",
            type: "Air Quality Index",
            typecode: 4,
            cost: 0.2,
            hoursConsumed: 6,
            created_date : "05/10/2016",
            status : "Active",
            end_date : "-"
        },
        {
            sensorid: "2114e2dqwdqg1",
            location: "Fresno",
            type: "Pressure",
            typecode: 2,
            cost: 0.2,
            hoursConsumed: 6,
            created_date : "05/10/2016",
            status : "Active",
            end_date : "-"
        },
        { sensorid: "2114e2dqwdqg2",
            location: "Fresno",
            type: "Temperature",
            typecode: 1,
            cost: 0.2,
            hoursConsumed: 6,
            created_date : "05/10/2016",
            status : "Active",
            end_date : "-"
        },
        { sensorid: "2114e2dqwdqg3",
            location: "Fresno",
            type: "Humidity",
            typecode: 3,
            cost: 0.2,
            hoursConsumed: 6,
            created_date : "05/10/2016",
            status : "Active",
            end_date : "-"
        },
        {sensorid: "2114e2dqwdqg4",
            location: "Fresno",
            type: "Air Quality Index",
            typecode: 4,
            cost: 0.2,
            hoursConsumed: 6,
            created_date : "05/10/2016",
            status : "Active",
            end_date : "-"
        },
        {
            sensorid: "2114e2dqwdqf1",
            location: "Apple Valley",
            type: "Pressure",
            typecode: 2,
            cost: 0.2,
            hoursConsumed: 6,
            created_date : "05/10/2016",
            status : "Active",
            end_date : "-"
        },
        {
            sensorid: "2114e2dqwdqf2",
            location: "Apple Valley",
            type: "Temperature",
            typecode: 1,
            cost: 0.2,
            hoursConsumed: 6,
            created_date : "05/10/2016",
            status : "Active",
            end_date : "-"
        },
        {
            sensorid: "2114e2dqwdqf3",
            location: "Apple Valley",
            type: "Humidity",
            typecode: 3,
            cost: 0.2,
            hoursConsumed: 6,
            created_date : "05/10/2016",
            status : "Active",
            end_date : "-"
        },
        {
            sensorid: "2114e2dqwdqf4",
            location: "Apple Valley",
            type: "Air Quality Index",
            typecode: 4,
            cost: 0.2,
            hoursConsumed: 6,
            created_date : "05/10/2016",
            status : "Active",
            end_date : "-"
        },{
            sensorid: "2114e2dqwdqk4",
            location: "Goleta",
            type: "Pressure",
            typecode: 2,
            cost: 0.2,
            hoursConsumed: 6,
            created_date : "05/10/2016",
            status : "Active",
            end_date : "-"
        },
        {
            sensorid: "2114e2dqwdqk3",
            location: "Goleta",
            type: "Temperature",
            typecode: 1,
            cost: 0.2,
            hoursConsumed: 6,
            created_date : "05/10/2016",
            status : "Active",
            end_date : "-"
        },
        {
            sensorid: "2114e2dqwdqk2",
            location: "Goleta",
            type: "Humidity",
            typecode: 3,
            cost: 0.2,
            hoursConsumed: 6,
            created_date : "05/10/2016",
            status : "Active",
            end_date : "-"
        },
        {
            sensorid: "2114e2dqwdqk1",
            location: "Goleta",
            type: "Air Quality Index",
            typecode: 4,
            cost: 0.2,
            hoursConsumed: 6,
            created_date : "05/10/2016",
            status : "Active",
            end_date : "-"
        }];



}]);
//.controller('MainCtrl',function($scope){

//})