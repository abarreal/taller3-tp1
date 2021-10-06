using Microsoft.AspNetCore.Mvc;
using Microsoft.AspNetCore.Mvc.Filters;
using Website.Visits;

namespace Website {
    public abstract class WebsiteController : Controller
    {
        private readonly VisitCounterService _visitcounterSvc;

        protected WebsiteController(VisitCounterService svc)
        {
            _visitcounterSvc = svc;
        }

        public override void OnActionExecuting(ActionExecutingContext context)
        {
            // Get the path of the request.
            var path = context.HttpContext.Request.Path;
            // Register the visit. Do it asynchronously, since there is eventual
            // consistency anyways.
            var registerVisitTask = _visitcounterSvc.CountVisitAsync(path);
            var retrieveCountTask = _visitcounterSvc.GetVisitCount();
            // We wait for the tasks to conclude.
            registerVisitTask.Wait();
            retrieveCountTask.Wait();
            // Set the count on the viewbag to be used on the view.
            ViewData["VisitCount"] = retrieveCountTask.Result;
        }

    }
}