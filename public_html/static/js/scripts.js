'use strict';
$(function () {

    graphLevels("levels");

    $('.mytab .nav-item').click(function(){
        $('.mytab .nav-item').removeClass('active');
        $(this).addClass('active');
        $('.tabs').addClass('d-none');
        
        $('.tab'+$(this).data("tab")).removeClass('d-none');
        return false;
    });
    
    

});