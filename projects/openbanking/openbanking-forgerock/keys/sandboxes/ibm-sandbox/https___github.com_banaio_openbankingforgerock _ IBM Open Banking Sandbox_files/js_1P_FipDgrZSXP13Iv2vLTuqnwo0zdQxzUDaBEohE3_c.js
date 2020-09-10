/********************************************************* {COPYRIGHT-TOP} ***
 * Licensed Materials - Property of IBM
 * 5725-L30, 5725-Z22
 *
 * (C) Copyright IBM Corporation 2016, 2019
 *
 * All Rights Reserved.
 * US Government Users Restricted Rights - Use, duplication or disclosure
 * restricted by GSA ADP Schedule Contract with IBM Corp.
 ********************************************************** {COPYRIGHT-END} **/

(function ($) {
    'use strict';

    function APP() {
        var self = this;

        var $window = $(window);

        function checkWidth() {
            var headerdiv = $("#header");
            if (headerdiv) {
                var headertop = headerdiv.offset().top;
                var scrolltop = $(document).scrollTop();
                var headerbottom = headertop - scrolltop + headerdiv.outerHeight(true);
                $("#columns").css({'margin-top': ((headerbottom) + 'px'), 'padding-top': '10px'});
                if ($("body").hasClass("page-comment-reply")) {
                    // move entire page over to allow comment content to appear at the top
                    $('#page').css({'margin-left': '180px'});
                } else {
                    // set toc top to height of header
                    $(".navigate-toc").css({
                        'top': ((headerbottom) + 'px'),
                        'height': ('calc(100% - ' + (headerbottom) + 'px)')
                    });
                }
            }

            var windowsize = $window.width();
            if (windowsize < 800) {
                // use hamburger menu
                $("#hamburger").removeClass("hidden");
                $('.navigate-toc').addClass('hidden');
            } else {
                $("#hamburger").addClass("hidden");
                $('.navigate-toc').removeClass('hidden');
            }
        }

        // Execute on load
        setTimeout(function () {
            checkWidth();
        }, 0);
        // Bind hamburger event listener
        $window.resize(checkWidth);

        $("#hamburger").click(function () {
            $('.navigate-toc').removeClass('hidden');
        });

        // close the menu
        $('.navigate-toc').click(function () {
            var windowsize = $window.width();
            if (windowsize < 800) {
                $('.navigate-toc').addClass('hidden');
            }
        });
        $('.mesh-portal-product').click(function () {
            var windowsize = $window.width();
            if (windowsize < 800) {
                $('.navigate-toc').addClass('hidden');
            }
        });
    }

    window.APP = APP;

    $(document).ready(function () {
        window.APP = new APP();
    });
}(jQuery));;
/**
 * @file
 * A JavaScript file for the theme.
 * This file should be used as a template for your other js files.
 * It defines a drupal behavior the "Drupal way".
 *
 */

// JavaScript should be made compatible with libraries other than jQuery by
// wrapping it with an "anonymous closure". See:
// - https://drupal.org/node/1446420
// - http://www.adequatelygood.com/2010/3/JavaScript-Module-Pattern-In-Depth

(function ($, Drupal, window, document, undefined) {
  'use strict';

  // To understand behaviors, see https://drupal.org/node/756722#behaviors
  Drupal.behaviors.hideSubmitBlockit = {
    attach: function(context) {
      var timeoutId = null;
      $('form', context).once('hideSubmitButton', function () {
        var $form = $(this);

        // Bind to input elements.
        if (Drupal.settings.hide_submit.hide_submit_method === 'indicator') {
          // Replace input elements with buttons.
          $('input.form-submit', $form).each(function(index, el) {
            var attrs = {};

            $.each($(this)[0].attributes, function(idx, attr) {
                attrs[attr.nodeName] = attr.nodeValue;
            });

            $(this).replaceWith(function() {
                return $("<button/>", attrs).append($(this).attr('value'));
            });
          });
          // Add needed attributes to the submit buttons.
          $('button.form-submit', $form).each(function(index, el) {
            $(this).addClass('ladda-button button').attr({
              'data-style': Drupal.settings.hide_submit.hide_submit_indicator_style,
              'data-spinner-color': Drupal.settings.hide_submit.hide_submit_spinner_color,
              'data-spinner-lines': Drupal.settings.hide_submit.hide_submit_spinner_lines
            });
          });
          Ladda.bind('.ladda-button', $form, {
            timeout: Drupal.settings.hide_submit.hide_submit_reset_time
          });
        }
        else {
          $('input.form-submit, button.form-submit', $form).click(function (e) {
            var el = $(this);
            el.after('<input type="hidden" name="' + el.attr('name') + '" value="' + el.attr('value') + '" />');
            return true;
          });
        }

        // Bind to form submit.
        $('form', context).submit(function (e) {
          var $inp;
          if (!e.isPropagationStopped()) {
            if (Drupal.settings.hide_submit.hide_submit_method === 'disable') {
              $('input.form-submit, button.form-submit', $form).attr('disabled', 'disabled').each(function (i) {
                var $button = $(this);
                if (Drupal.settings.hide_submit.hide_submit_css) {
                  $button.addClass(Drupal.settings.hide_submit.hide_submit_css);
                }
                if (Drupal.settings.hide_submit.hide_submit_abtext) {
                  $button.val($button.val() + ' ' + Drupal.settings.hide_submit.hide_submit_abtext);
                }
                $inp = $button;
              });

              if ($inp && Drupal.settings.hide_submit.hide_submit_atext) {
                $inp.after('<span class="hide-submit-text">' + Drupal.checkPlain(Drupal.settings.hide_submit.hide_submit_atext) + '</span>');
              }
            }
            else if (Drupal.settings.hide_submit.hide_submit_method !== 'indicator'){
              var pdiv = '<div class="hide-submit-text' + (Drupal.settings.hide_submit.hide_submit_hide_css ? ' ' + Drupal.checkPlain(Drupal.settings.hide_submit.hide_submit_hide_css) + '"' : '') + '>' + Drupal.checkPlain(Drupal.settings.hide_submit.hide_submit_hide_text) + '</div>';
              if (Drupal.settings.hide_submit.hide_submit_hide_fx) {
                $('input.form-submit, button.form-submit', $form).addClass(Drupal.settings.hide_submit.hide_submit_css).fadeOut(100).eq(0).after(pdiv);
                $('input.form-submit, button.form-submit', $form).next().fadeIn(100);
              }
              else {
                $('input.form-submit, button.form-submit', $form).addClass(Drupal.settings.hide_submit.hide_submit_css).hide().eq(0).after(pdiv);
              }
            }
            // Add a timeout to reset the buttons (if needed).
            if (Drupal.settings.hide_submit.hide_submit_reset_time) {
              timeoutId = window.setTimeout(function() {
                hideSubmitResetButtons(null, $form);
              }, Drupal.settings.hide_submit.hide_submit_reset_time);
            }
          }
          return true;
        });
      });

      // Bind to clientsideValidationFormHasErrors to support clientside validation.
      // $(document).bind('clientsideValidationFormHasErrors', function(event, form) {
        //hideSubmitResetButtons(event, form.form);
      // });

      // Reset all buttons.
      function hideSubmitResetButtons(event, form) {
        // Clear timer.
        window.clearTimeout(timeoutId);
        timeoutId = null;
        switch (Drupal.settings.hide_submit.hide_submit_method) {
          case 'disable':
            $('input.' + Drupal.checkPlain(Drupal.settings.hide_submit.hide_submit_css) + ', button.' + Drupal.checkPlain(Drupal.settings.hide_submit.hide_submit_css), form)
              .each(function (i, el) {
                $(el).removeClass(Drupal.checkPlain(Drupal.settings.hide_submit.hide_submit_hide_css))
                  .removeAttr('disabled');
              });
            $('.hide-submit-text', form).remove();
            break;

          case 'indicator':
            Ladda.stopAll();
            break;

          default:
            $('input.' + Drupal.checkPlain(Drupal.settings.hide_submit.hide_submit_css) + ', button.' + Drupal.checkPlain(Drupal.settings.hide_submit.hide_submit_css), form)
              .each(function (i, el) {
                $(el).stop()
                  .removeClass(Drupal.checkPlain(Drupal.settings.hide_submit.hide_submit_hide_css))
                  .show();
              });
            $('.hide-submit-text', form).remove();
        }
      }
    }
  };

})(jQuery, Drupal, window, this.document);
;
