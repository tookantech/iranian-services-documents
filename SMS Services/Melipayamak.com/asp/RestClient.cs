using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Collections.Specialized;
using System.IO;
using System.Linq;
using System.Net;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;

namespace MPMVC.Services
{
	public class RestClient
	{
		private const string endpoint = "https://rest.payamak-panel.com/api/SendSMS/";
		private const string sendOp = "SendSMS";

		private const string getDeliveryOp = "GetDeliveries2";
		private const string getMessagesOp = "GetMessages";
		private const string getCreditOp = "GetCredit";
		private const string getBasePriceOp = "GetBasePrice";
		private const string getUserNumbersOp = "GetUserNumbers";
		private const string sendByBaseNumberOp = "BaseServiceNumber";

		private string UserName;
		private string Password;

		public RestClient(string username, string password)
		{
			UserName = username;
			Password = password;
		}

		private RestResponse makeRequest(Dictionary<string, string> values, string op)
		{
			//HttpWebRequest request = WebRequest.CreateHttp(endpoint + op);
			//request.Method = "POST";

			//using (HttpWebResponse response = (HttpWebResponse)request.GetResponse())
			//{
			//	using (Stream responseStream = response.GetResponseStream())
			//	{
			//		using (StreamReader myStreamReader = new StreamReader(responseStream, Encoding.UTF8))
			//		{
			//			string responseJSON = myStreamReader.ReadToEnd();
			//			return JsonConvert.DeserializeObject<RestResponse>(responseJSON);
			//		}
			//	}
			//}

			using (var client = new WebClient())
			{
				var response = client.UploadValues(endpoint + op, values.Aggregate(new NameValueCollection(),
				(seed, current) => {
					seed.Add(current.Key, current.Value);
					return seed;
				}));

				return JsonConvert.DeserializeObject<RestResponse>(Encoding.Default.GetString(response));
			}
		}

		



		public RestResponse Send(string to, string from, string message, bool isflash)
		{
			var values = new Dictionary<string, string>
			{
				{ "username", UserName },
				{ "password", Password },
				{ "to" , to },
				{ "from" , from },
				{ "text" , message },
				{ "isFlash" , isflash.ToString() }
			};

			return makeRequest(values, sendOp);
		}

		public RestResponse SendByBaseNumber(string text, string to, int bodyId)
		{
			var values = new Dictionary<string, string>
			{
				{ "username", UserName },
				{ "password", Password },
				{ "text" , text },
				{ "to" , to },
				{ "bodyId" , bodyId.ToString() }
			};

			return makeRequest(values, sendByBaseNumberOp);
		}

		public RestResponse GetDelivery(long recid)
		{
			var values = new Dictionary<string, string>
			{
				{ "UserName", UserName },
				{ "PassWord", Password },
				{ "recID" , recid.ToString() },
			};

			return makeRequest(values, getDeliveryOp);
		}


		public RestResponse GetMessages(int location, string from, string index, int count)
		{
			var values = new Dictionary<string, string>
			{
				{ "UserName", UserName },
				{ "PassWord", Password },
				{ "Location" , location.ToString() },
				{ "From" , from },
				{ "Index" , index },
				{ "Count" , count.ToString() }
			};

			return makeRequest(values, getMessagesOp);
		}

		public RestResponse GetCredit()
		{
			var values = new Dictionary<string, string>
			{
				{ "UserName", UserName },
				{ "PassWord", Password },
			};

			return makeRequest(values, getCreditOp);
		}

		public RestResponse GetBasePrice()
		{
			var values = new Dictionary<string, string>
			{
				{ "UserName", UserName },
				{ "PassWord", Password },
			};

			return makeRequest(values, getBasePriceOp);
		}

		public RestResponse GetUserNumbers()
		{
			var values = new Dictionary<string, string>
			{
				{ "UserName", UserName },
				{ "PassWord", Password },
			};

			return makeRequest(values, getUserNumbersOp);
		}

	}
	
}
