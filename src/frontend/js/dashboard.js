/**
 * Created by shikher on 14-Apr-16.
 */
(function () {

    userId = 1234

    var app = angular.module('sensorDataModule');

    app.controller('sensorDataController', ['$http', '$scope', function ($http, $scope) {
        // Set up the chart
        $scope.tempchart = new Highcharts.Chart({
            chart: {
                renderTo: 'temp-level-container',
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
            title: {
                text: 'Temperature Sensor Data'
            },
            subtitle: {
                text: 'Measured in celsius'
            },
            plotOptions: {
                column: {
                    depth: 30
                }

            },
            series: [{
                data: {},
                name: "Temperature"
            }]
        });


        $scope.waterchart = new Highcharts.Chart({
            chart: {
                renderTo: 'water-level-container',
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
            title: {
                text: 'Water Level Sensor Data'
            },
            subtitle: {
                text: 'Measured in feet'
            },
            plotOptions: {
                column: {
                    depth: 30
                }

            },
            series: [{
                data: {},
                name: "Water Level"
            }]
        });

        $scope.getTempData = function () {
            $http.get().success(function (data) {
                $scope.tempchart.series.data = data.value;
            });
        }

        $scope.getWaterLevelData = function () {
            $http.get().success(function (data) {
                $scope.waterchart.series.data = data.value;
            });
        }
    }]);

})();