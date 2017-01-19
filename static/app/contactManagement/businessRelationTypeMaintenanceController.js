/**
 * Created by tuanha-01 on 5/13/2016.
 */
"use strict";

define(['angularAMD', 'ajaxService', 'alertsService', 'businessRelationTypesService'], function (angularAMD, $) {
    var injectParams = ['$scope', '$rootScope', '$state', '$window', 'moment', 'alertsService', 'businessRelationTypesService', '$stateParams', 'Constants'];

    var businessRelationTypeMaintenanceController = function ($scope, $rootScope, $state, $window, moment, alertsService, businessRelationTypesService, $stateParams, Constants) {

        $scope.initializeController = function () {
            $rootScope.applicationModule = "BusinessRelationTypeMaintenance";
            $rootScope.alerts = [];

            var businessRelationTypeID = ($stateParams.id || null);
            
            $scope.ID = businessRelationTypeID;
                        
            $scope.Constants = Constants;

            if (businessRelationTypeID == null) {
                $scope.Code = "";
                $scope.Description = "";
                $scope.Status = $scope.Constants.Status[1].Code;
                $scope.RecCreatedByID = $rootScope.currentUser.ID;
                $scope.RecCreatedByUser = $rootScope.currentUser.Name;
                $scope.RecCreated = new Date();
                $scope.RecModifiedByID = $rootScope.currentUser.ID;
                $scope.RecModifiedByUser = $rootScope.currentUser.Name;
                $scope.RecModified = new Date();
            } else {
                var getBusinessRelationType = new Object();
                getBusinessRelationType.ID = businessRelationTypeID
                businessRelationTypesService.getBusinessRelationType(getBusinessRelationType, $scope.businessRelationTypeCompleted, $scope.businessRelationTypeError);
            }
        };

        $scope.businessRelationTypeCompleted = function (response, status) {
            $scope.ID = response.Data.BusinessRelationType.ID;
            $scope.Code = response.Data.BusinessRelationType.Code;
            $scope.Description = response.Data.BusinessRelationType.Description;
            $scope.Status = response.Data.BusinessRelationType.Status;
            $scope.Version = response.Data.BusinessRelationType.Version;
            $scope.ClientID = response.Data.BusinessRelationType.ClientID;
            $scope.OrganizationID = response.Data.BusinessRelationType.OrganizationID;
            $scope.RecCreatedByID = response.Data.BusinessRelationType.RecCreatedByID;
            $scope.RecCreatedByUser = response.Data.BusinessRelationType.RecCreatedByUser;
            $scope.RecCreated = new moment.unix(response.Data.BusinessRelationType.RecCreated).toDate();
            $scope.RecModifiedByID = response.Data.BusinessRelationType.RecModifiedByID;
            $scope.RecModifiedByUser = response.Data.BusinessRelationType.RecModifiedByUser;
            $scope.RecModified = new moment.unix(response.Data.BusinessRelationType.RecModified).toDate();
        };

        $scope.businessRelationTypeError = function (response, status) {
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
                            Table: "business_relation_type",
                            //RecID: $scope.ID 
                            //Don't use. $scope.ID = "" when set validationOptions. 
                            //Use function to get current $scope.ID
                            RecID: function() { 
                                return $scope.ID 
                            }
                        }
                    }
                },
                Description: {
                    required: true
                }
            }
        };

        $scope.update = function (form) {
            if(form.validate()) {
                var businessRelationType = $scope.createBusinessRelationTypeObject();
                businessRelationTypesService.updateBusinessRelationType(businessRelationType, $scope.businessRelationTypeUpdateCompleted, $scope.businessRelationTypeUpdateError)
            }
        };

        $scope.cancel = function (form) {
           setTimeout(function() {
                $state.go('businessRelationType', { businessRelationTypeID : $scope.ID });
            }, 10);
        };

        $scope.businessRelationTypeUpdateCompleted = function (response, status) {
            $scope.ID = response.Data.BusinessRelationType.ID;
            alertsService.RenderSuccessMessage(response.ReturnMessage);
            
            setTimeout(function() {
                $state.go('businessRelationType', { businessRelationTypeID : $scope.ID });
            }, 1000);
        };

        $scope.businessRelationTypeUpdateError = function (response, status) {
            alertsService.RenderErrorMessage(response.Error);
        }

        $scope.createBusinessRelationTypeObject = function () {
            var businessRelationType = new Object();
            businessRelationType.ID = $scope.ID;
            businessRelationType.Code = $scope.Code;
            businessRelationType.Description = $scope.Description;
            businessRelationType.Version = $scope.Version;

            businessRelationType.Status = $scope.Status;
            businessRelationType.ClientID = $scope.ClientID;
            businessRelationType.OrganizationID = $scope.OrganizationID;
            businessRelationType.RecCreatedByID = $scope.RecCreatedByID;
            businessRelationType.RecCreatedByUser = $scope.RecCreatedByUser;
            businessRelationType.RecCreated = new moment($scope.RecCreated).unix();
            businessRelationType.RecModifiedByID = $rootScope.currentUser.ID;
            businessRelationType.RecModifiedByUser = $rootScope.currentUser.Name;
            businessRelationType.RecModified = new moment($scope.RecModified).unix();

            return businessRelationType;
        }
    };

    businessRelationTypeMaintenanceController.$inject = injectParams;
    angularAMD.controller('BusinessRelationTypeMaintenanceController', businessRelationTypeMaintenanceController);
});
