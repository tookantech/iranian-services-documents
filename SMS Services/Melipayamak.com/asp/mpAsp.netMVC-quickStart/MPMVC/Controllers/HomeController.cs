using MPMVC.ServiceReference;
using MPMVC.Services;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using System.Web;
using System.Web.Mvc;

namespace MPMVC.Controllers
{
	public class HomeController : Controller
	{
		public ActionResult Index()
		{
			return View();
		}

		public ActionResult About()
		{
			ViewBag.Message = "Your application description page.";

			return View();
		}

		public ActionResult Contact()
		{
			ViewBag.Message = "Your contact page.";

			return View();
		}


		public ActionResult Soap()
		{
			const string username = "username";
			const string password = "password";
			const string from = "5000...";
			const string to = "09123456789";
			const string text = "تست وب سرویس ملی پیامک";
			const bool isFlash = false;
			SendSoapClient soapClient = new SendSoapClient();
			var result = soapClient.SendSimpleSMS2(username, password, to, from, text, isFlash);
			return Json(result, JsonRequestBehavior.AllowGet);
		}

		public async Task<ActionResult> SoapAsync()
		{
			const string username = "username";
			const string password = "password";
			const string from = "5000...";
			const string to = "09123456789";
			const string text = "تست وب سرویس ملی پیامک";
			const bool isFlash = false;
			SendSoapClient soapClient = new SendSoapClient();
			var result = await soapClient.SendSimpleSMS2Async(username, password, to, from, text, isFlash);
			return Json(result, JsonRequestBehavior.AllowGet);
		}


		public ActionResult Rest()
		{
			const string username = "username";
			const string password = "password";
			const string from = "5000...";
			const string to = "09123456789";
			const string text = "تست وب سرویس ملی پیامک";
			const bool isFlash = false;
			RestClient restClient = new RestClient(username, password);
			var result = restClient.Send(to, from, text, isFlash);
			return Json(result, JsonRequestBehavior.AllowGet);
		}

		public async Task<ActionResult> RestAsync()
		{
			const string username = "username";
			const string password =  "password";
			const string from = "5000...";
			const string to = "09123456789";
			const string text = "تست وب سرویس ملی پیامک";
			const bool isFlash = false;
			RestClientAsync restClient = new RestClientAsync(username, password);
			var result = await restClient.SendAsync(to, from, text, isFlash);
			return Json(result, JsonRequestBehavior.AllowGet);
		}
	}
}