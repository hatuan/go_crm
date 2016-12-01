/**
 * Created by tuanha-01 on 5/30/2016.
 */
define(['angularAMD', 'ajaxService'], function (angularAMD) {

    angularAMD.service('businessRelationSectorsService', ['ajaxService', function (ajaxService) {
        
        this.getBusinessRelationSectors = function (businessRelationSector, successFunction, errorFunction) {
            ajaxService.AjaxGetWithData(businessRelationSector, "/api/businessrelationsectors", successFunction, errorFunction);
        };

        this.updateBusinessRelationSector = function (businessRelationSector, successFunction, errorFunction) {
            ajaxService.AjaxPost(businessRelationSector, "/api/businessrelationsectors", successFunction, errorFunction);
        };

        this.deleteBusinessRelationSector = function (data, successFunction, errorFunction) {
            ajaxService.AjaxDelete(data, "/api/businessrelationsectors", successFunction, errorFunction);
        };

        this.getBusinessRelationSector = function (businessRelationSector, successFunction, errorFunction) {
            ajaxService.AjaxGetWithData(businessRelationSector, "/api/businessrelationsector", successFunction, errorFunction);
        };    
    }]);
});
