/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['angularAMD', 'ajaxService', 'alertsService', 'businessRelationSectorsService'], function (angularAMD, $) {
    var injectParams = ['$scope', '$rootScope', '$state', '$window', 'moment', 'alertsService', 'businessRelationSectorsService', '$stateParams', 'Constants'];

    var businessRelationSectorMaintenanceController = function ($scope, $rootScope, $state, $window, moment, alertsService, businessRelationSectorsService, $stateParams, Constants) {

        $scope.initializeController = function () {
            $rootScope.applicationModule = "BusinessRelationSectorMaintenance";
            $rootScope.alerts = [];

            var businessRelationSectorID = ($stateParams.id || "");
            
            $scope.ID = businessRelationSectorID;
                        
            $scope.Constants = Constants;

            if (businessRelationSectorID == "") {
                $scope.ID = "";
                $scope.Code = "";
                $scope.Name = "";
                $scope.Status = $scope.Constants.Status[1].Code;
                $scope.ClientID = "";
                $scope.OrganizationID = "";
                $scope.RecCreatedByID = $rootScope.currentUser.ID;
                $scope.RecCreatedByUser = $rootScope.currentUser.Name;
                $scope.RecCreated = new Date();
                $scope.RecModifiedByID = $rootScope.currentUser.ID;
                $scope.RecModifiedByUser = $rootScope.currentUser.Name;
                $scope.RecModified = new Date();
            } else {
                var getBusinessRelationSector = new Object();
                getBusinessRelationSector.ID = businessRelationSectorID
                businessRelationSectorsService.getBusinessRelationSector(getBusinessRelationSector, $scope.businessRelationSectorCompleted, $scope.businessRelationSectorError);
            }
        };

        $scope.businessRelationSectorCompleted = function (response, status) {
            $scope.ID = response.Data.BusinessRelationSector.ID;
            $scope.Code = response.Data.BusinessRelationSector.Code;
            $scope.Name = response.Data.BusinessRelationSector.Name;
            $scope.Status = response.Data.BusinessRelationSector.Status;
            $scope.Version = response.Data.BusinessRelationSector.Version;
            $scope.ClientID = response.Data.BusinessRelationSector.ClientID;
            $scope.OrganizationID = response.Data.BusinessRelationSector.OrganizationID;
            $scope.RecCreatedByID = response.Data.BusinessRelationSector.RecCreatedByID;
            $scope.RecCreatedByUser = response.Data.BusinessRelationSector.RecCreatedByUser;
            $scope.RecCreated = new moment.unix(response.Data.BusinessRelationSector.RecCreated).toDate();
            $scope.RecModifiedByID = response.Data.BusinessRelationSector.RecModifiedByID;
            $scope.RecModifiedByUser = response.Data.BusinessRelationSector.RecModifiedByUser;
            $scope.RecModified = new moment.unix(response.Data.BusinessRelationSector.RecModified).toDate();
        };

        $scope.businessRelationSectorError = function (response, status) {
            alertsService.RenderErrorMessage(response.Error);
        }

        $scope.validationOptions = {
            rules: {
                Code: {
                    required: true,
                    "remote": {
                        url: "api/check-unique",
                        type: "post",
                        //dataType: 'json', //dataType is json but don't have any effect. 
                        data: {
                            UserID : function() {
                                return $rootScope.currentUser.ID
                            }, 
                            Table: "business_relation_sector",
                            //RecID: $scope.ID 
                            //Don't use. $scope.ID = "" when set validationOptions. 
                            //Use function to get current $scope.ID
                            RecID: function() { 
                                return $scope.ID 
                            }
                        }
                    }
                },
                Name: {
                    required: true
                }
            }
        };

        $scope.update = function (form) {
            if(form.validate()) {
                var businessRelationSector = $scope.createBusinessRelationSectorObject();
                businessRelationSectorsService.updateBusinessRelationSector(businessRelationSector, $scope.businessRelationSectorUpdateCompleted, $scope.businessRelationSectorUpdateError)
            }
        };

        $scope.cancel = function (form) {
           setTimeout(function() {
                $state.go('businessRelationSector', { businessRelationSectorID : $scope.ID });
            }, 10);
        };

        $scope.businessRelationSectorUpdateCompleted = function (response, status) {
            $scope.ID = response.Data.BusinessRelationSector.ID;
            alertsService.RenderSuccessMessage(response.ReturnMessage);
            
            setTimeout(function() {
                $state.go('businessRelationSector', { businessRelationSectorID : $scope.ID });
            }, 1000);
        };

        $scope.businessRelationSectorUpdateError = function (response, status) {
            alertsService.RenderErrorMessage(response.Error);
        }

        $scope.createBusinessRelationSectorObject = function () {
            var businessRelationSector = new Object();
            businessRelationSector.ID = $scope.ID;
            businessRelationSector.Code = $scope.Code;
            businessRelationSector.Name = $scope.Name;
            businessRelationSector.Version = $scope.Version;

            businessRelationSector.Status = $scope.Status;
            businessRelationSector.ClientID = $scope.ClientID;
            businessRelationSector.OrganizationID = $scope.OrganizationID;
            businessRelationSector.RecCreatedByID = $scope.RecCreatedByID;
            businessRelationSector.RecCreatedByUser = $scope.RecCreatedByUser;
            businessRelationSector.RecCreated = new moment($scope.RecCreated).unix();
            businessRelationSector.RecModifiedByID = $rootScope.currentUser.ID;
            businessRelationSector.RecModifiedByUser = $rootScope.currentUser.Name;
            businessRelationSector.RecModified = new moment($scope.RecModified).unix();

            return businessRelationSector;
        }
    };

    businessRelationSectorMaintenanceController.$inject = injectParams;
    angularAMD.controller('BusinessRelationSectorMaintenanceController', businessRelationSectorMaintenanceController);
});
