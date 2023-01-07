using Newtonsoft.Json;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Net.Http;
using System.Text;
using System.Threading.Tasks;

namespace MPMVC.Services
{
	public class RestClientAsync
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

		public RestClientAsync(string username, string password)
		{
			UserName = username;
			Password = password;
		}

		private async Task<RestResponse> makeRequestAsync(Dictionary<string, string> values, string op)
		{
			var content = new FormUrlEncodedContent(values);

			using (var httpClient = new HttpClient())
			{
				var response = await httpClient.PostAsync(endpoint + op, content);
				var responseString = await response.Content.ReadAsStringAsync();

				return JsonConvert.DeserializeObject<RestResponse>(responseString);
			}
		}



		public async Task<RestResponse> SendAsync(string to, string from, string message, bool isflash)
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

			return await makeRequestAsync(values, sendOp);
		}

		public async Task<RestResponse> SendByBaseNumberAsync(string text, string to, int bodyId)
		{
			var values = new Dictionary<string, string>
			{
				{ "username", UserName },
				{ "password", Password },
				{ "text" , text },
				{ "to" , to },
				{ "bodyId" , bodyId.ToString() }
			};

			return await makeRequestAsync(values, sendByBaseNumberOp);
		}

		public async Task<RestResponse> GetDeliveryAsync(long recid)
		{
			var values = new Dictionary<string, string>
			{
				{ "UserName", UserName },
				{ "PassWord", Password },
				{ "recID" , recid.ToString() },
			};

			return await makeRequestAsync(values, getDeliveryOp);
		}


		public async Task<RestResponse> GetMessagesAsync(int location, string from, string index, int count)
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

			return await makeRequestAsync(values, getMessagesOp);
		}

		public async Task<RestResponse> GetCreditAsync()
		{
			var values = new Dictionary<string, string>
			{
				{ "UserName", UserName },
				{ "PassWord", Password },
			};

			return await makeRequestAsync(values, getCreditOp);
		}

		public async Task<RestResponse> GetBasePriceAsync()
		{
			var values = new Dictionary<string, string>
			{
				{ "UserName", UserName },
				{ "PassWord", Password },
			};

			return await makeRequestAsync(values, getBasePriceOp);
		}

		public async Task<RestResponse> GetUserNumbersAsync()
		{
			var values = new Dictionary<string, string>
			{
				{ "UserName", UserName },
				{ "PassWord", Password },
			};

			return await makeRequestAsync(values, getUserNumbersOp);
		}

	}


	//response class
	public class RestResponse
	{
		public string Value { get; set; }
		public int RetStatus { get; set; }
		public string StrRetStatus { get; set; }
	}
}
