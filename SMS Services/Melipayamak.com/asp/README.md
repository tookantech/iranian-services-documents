# melipayamak Asp.net MVC

<div dir='rtl'>

### معرفی وب سرویس ملی پیامک
ملی پیامک یک وب سرویس کامل برای ارسال و دریافت پیامک و پیامک صوتی و مدیریت کامل خدمات دیگر است که براحتی میتوانید از آن استفاده کنید.

<hr>

### نصب

<p>قبل از نصب نیاز به ثبت نام در سایت ملی پیامک دارید.</p>

[**ثبت نام به همراه دریافت 200 پیامک هدیه جهت تست وبسرویس**](http://www.melipayamak.com/)

پس از ثبت نام، وب سرویس‌های ملی پیامک را از طریق آدرسهای زیر به عنوان **Service Reference** به پروژه خود اضافه کنید.

</div>

<div dir='rtl'>
  
#### وب سرویس ارسال پیام

</div>

> [api.payamak-panel.com/post/Send.asmx](http://api.payamak-panel.com/post/Send.asmx)


<div dir='rtl'>
  
#### وب سرویس دریافت پیام

</div>

> [api.payamak-panel.com/post/receive.asmx](http://api.payamak-panel.com/post/receive.asmx)


<div dir='rtl'>
  
#### وب سرویس مدیریت مخاطبین

</div>

> [api.payamak-panel.com/post/contacts.asmx](http://api.payamak-panel.com/post/contacts.asmx)

> [api.payamak-panel.com/post/Actions.asmx](http://api.payamak-panel.com/post/Actions.asmx)

<div dir='rtl'>
  
#### وب سرویس ارسال زماندار

</div>

> [api.payamak-panel.com/post/Schedule.asmx](http://api.payamak-panel.com/post/Schedule.asmx)


<div dir='rtl'>
  
#### وب سرویس پشتیبانی کاربران

</div>

> [api.payamak-panel.com/post/Tickets.asmx](http://api.payamak-panel.com/post/Tickets.asmx)


<div dir='rtl'>
  
#### وب سرویس مدیریت کاربران

</div>

> [api.payamak-panel.com/post/Users.asmx](http://api.payamak-panel.com/post/Users.asmx)


<div dir='rtl'>
  
#### وب سرویس ارسال صوتی

</div>

> [api.payamak-panel.com/post/Voice.asmx](http://api.payamak-panel.com/post/Voice.asmx)


<div dir='rtl'>
  
شما می توانید از طریق [Nuget](https://www.nuget.org/packages/Melipayamak.RestClient/) نیز وب سرویس ملی پیامک را به پروزه خود اضافه نمایید :
  
</div>

```
Install-Package Melipayamak.RestClient -Version 1.0.0
```

<div dir='rtl'>
  
#### نحوه استفاده

نمونه کد سی شارپ برای ارسال پیامک به صورت زیر است که شما می توانید در پیکرۀ Controller های برنامه خود از آنها استفاده کنید
شما می توانید به طور دلخواه از فرم Asynchron این متدها هرکجا لازم شد استفاده کنید

</div>



```js
const string username = "username";
const string password = "password";
const string from = "5000...";
const string to = "09123456789";
const string text = "تست وب سرویس ملی پیامک";
const bool isFlash = false;
SendSoapClient soapClient = new SendSoapClient();
soapClient.SendSimpleSMS2(username, password, to, from, text, isFlash);
//یا برای ارسال به مجموعه ای از مخاطبین
soapClient.SendSimpleSMS(username, password, new string[] { to }, from, text, isFlash);
```

<div dir='rtl'>

از آنجا که وب سرویس ملی پیامک تنها محدود به ارسال پیامک نیست شما از طریق زیر میتوانید به وب سرویس ها دسترسی کامل داشته باشید:
</div>

```js
// وب سرویس پیامک
RestClient restClient = new RestClient(username, password);
SendSoapClient soapClient = new SendSoapClient();
// وب سرویس تیکت پشتیبانی
TicketsSoapClient ticketSoapClient = new TicketsSoapClient();
// وب سرویس برای مدیریت کامل  ارسال انبوه پیامک
ActionsSoapClient actionSoapClient = new ActionsSoapClient();
//وب سرویس کاربران
UsersSoapClient usersSoapClient = new UsersSoapClient();
//وب سرویس دفترچه تلفن
ContactsSoapClient contactSoapClient = new ContactsSoapClient();
//وب سرویس زمان بندی
ScheduleSoapClient scheduleSoapClient = new ScheduleSoapClient();
//وب سرویس پیام صوتی
VoiceSoapClient voiceSoapClient = new VoiceSoapClient();
//وب سرویس دریافت
ReceiveSoapClient receiveSoapClient = new ReceiveSoapClient();
```

<div dir='rtl'>
  
#### تفاوت های وب سرویس پیامک rest و soap

از آنجا که ملی پیامک وب سرویس کاملی رو در اختیار توسعه دهندگان میگزارد برای راحتی کار با وب سرویس پیامک علاوه بر وب سرویس اصلی soap وب سرویس rest رو هم در اختیار توسعه دهندگان گزاشته شده تا راحتتر بتوانند با وب سرویس کار کنند. تفاوت اصلی این دو در تعداد متد هاییست که میتوانید با آن کار کنید. برای کار های پایه میتوان از وب سرویس rest استفاده کرد برای دسترسی بیشتر و استفاده پیشرفته تر نیز باید از وب سرویس باید از وب سرویس soap استفاده کرد. جهت مطالعه بیشتر وب سرویس ها به قسمت وب سرویس پنل خود مراجعه کنید.

<hr/>


###  اطلاعات بیشتر

برای مطالعه بیشتر و دریافت راهنمای وب سرویس ها و آشنایی با پارامتر های ورودی و خروجی وب سرویس به صفحه معرفی
[وب سرویس ملی پیامک](https://github.com/Melipayamak/Webservices)
مراجعه نمایید .


<hr/>


### وب سرویس پیامک

متد های وب سرویس:

</div>

#### ارسال

```js
restClient.Send(to, from, text, isFlash);
soapClient.SendSimpleSMS(username, password, new string[] { to }, from, text, isFlash);
```
<div dir='rtl'>
  در آرگومان سوم روش soap میتوانید از هر تعداد مخاطب به عنوان آرایه استفاده کنید
</div>

#### ارسال از خط خدماتی اشتراکی

```js
restClient.SendByBaseNumber(text, to, bodyId);
soapClient.SendByBaseNumber2(username, password, text, to, bodyId);
```

#### دریافت وضعیت ارسال
```js
restClient.GetDelivery(recId);
soapClient.GetDelivery(recId);
soapClient.GetDeliveries(recIds[], username, password);
```

#### لیست پیامک ها

```js
restClient.GetMessages(location, index, count, from);
soapClient.getMessages(username, password, location, from, index, count);
// جهت دریافت به صورت رشته ای
receiveSoapClient.GetMessagesByDate(username, password, location, from, index, count, dateFrom, dateTo);
//جهت دریافت بر اساس تاریخ
receiveSoapClient.GetUsersMessagesByDate(username, password, location, from, index, count, dateFrom, dateTo);
// جهت دریافت پیام های کاربران بر اساس تاریخ 
```

#### موجودی
```js
restClient.GetCredit();
soapClient.GetCredit(username, password);
```

#### تعرفه پایه / دریافت قیمت قبل از ارسال
```js
restClient.GetBasePrice();
soapClient.GetSmsPrice(username, password, irancellCount, mtnCount, from, text);
```
#### لیست شماره اختصاصی
```js
usersSoapClient.GetUserNumbers(username, password);
```

#### بررسی تعداد پیامک های دریافتی
```js
soapClient.GetInboxCount(username, password, isRead);
//پیش فرض خوانده نشده 
```

#### ارسال پیامک پیشرفته
```js
soapClient.SendSms(username, password, to[], from, text, isflash, udh, recId[], status[]);
```

#### مشاهده مشخصات پیام
```js
receiveSoapClient.GetMessagesReceptions(username, password, msgId, fromRows);
```


#### حذف پیام دریافتی
```js
receiveSoapClient.RemoveMessages2(username, password, location, msgIds);
```


#### ارسال زماندار
```js
scheduleSoapClient.AddSchedule(username, password, to, from, text, isflash, scheduleDateTime, period);
```

#### ارسال زماندار متناظر
```js
scheduleSoapClient.AddMultipleSchedule(username, password, to[], from, text[], isflash, scheduleDateTime[], period);
```


#### ارسال سررسید
```js
scheduleSoapClient.AddNewUsance(username, password, to, from, text, isflash, scheduleStartDateTime, countRepeat, scheduleEndDateTime, periodType);
```

#### مشاهده وضعیت ارسال زماندار
```js
scheduleSoapClient.GetScheduleStatus(username, password, schId);
```

#### حذف پیامک زماندار
```js
scheduleSoapClient.RemoveSchedule(username, password, schId);
```


<div dir='rtl'>

### وب سرویس پیامک صوتی

</div>

####  ارسال پیامک همراه با تماس صوتی
```js
voiceSoapClient.SendSMSWithSpeechText(username, password, smsBody, speechBody, from, to);
```

####  ارسال پیامک همراه با تماس صوتی به صورت زمانبندی
```js
voiceSoapClient.SendSMSWithSpeechTextBySchduleDate(username, password, smsBody, speechBody, from, to, scheduleDate);
```

####  دریافت وضعیت پیامک همراه با تماس صوتی 
```js
voiceSoapClient.GetSendSMSWithSpeechTextStatus(username, password, recId);
```

#### تماس انبوه زماندار
```js
voiceSoapClient.SendBulkSpeechText(username, password, title, body, receivers, DateToSend, repeatCount);
```

#### تماس انبوه زماندار با انتخاب فایل
```js
voiceSoapClient.SendBulkVoiceSMS(username, password, title, fileId, receivers, DateToSend, repeatCount);
```

#### آپلود فایل صوتی
```js
voiceSoapClient.UploadVoiceFile(username, password, title, file);
```

<div dir='rtl'>
  
### وب سرویس ارسال انبوه/منطقه ای

</div>

#### دریافت شناسه شاخه های بانک شماره
```js
actionSoapClient.GetBranchs(username, password, owner);
```


#### اضافه کردن یک بانک شماره جدید
```js
actionSoapClient.AddBranch(username, password, branchName, owner);
```

#### اضافه کردن شماره به بانک
```js
actionSoapClient.AddNumber(username, password, branchId, mobileNumbers[]);
```

#### حذف یک بانک
```js
actionSoapClient.RemoveBranch(username, password, branchId);
```

#### ارسال انبوه از طریق بانک
```js
actionSoapClient.AddBulk(username, password, from, branch, bulkType, title, message, rangeFrom, rangeTo, DateToSend, requestCount, rowFrom);
```

#### تعداد شماره های موجود
```js
actionSoapClient.GetBulkCount(username, password, branch, rangeFrom, rangeTo);
```

#### گزارش گیری از ارسال انبوه
```js
actionSoapClient.GetBulkReceptions(username, password, bulkId, fromRows);
```


#### تعیین وضعیت ارسال 
```js
actionSoapClient.GetBulkStatus(username, password, bulkId, sent, failed, status);
```

#### تعداد ارسال های امروز
```js
actionSoapClient.GetTodaySent(username, password);
```

#### تعداد ارسال های کل

```js
actionSoapClient.GetTotalSent(username, password);
```

#### حذف ارسال منطقه ای
```js
actionSoapClient.RemoveBulk(username, password, bulkId);
```

#### ارسال متناظر
```js
actionSoapClient.SendMultipleSMS(username, password, to[], from, text[], isflash, udh, recId[], status);
```

#### نمایش دهنده وضعیت گزارش گیری

```js
actionSoapClient.UpdateBulkDelivery(username, password, bulkId);
```
<div dir='rtl'>
  
### وب سرویس تیکت

</div>

#### ثبت تیکت جدید
```js
ticketSoapClient.AddTicket(username, password, title, content, aletWithSms);
```

#### جستجو و دریافت تیکت ها

```js
ticketSoapClient.GetReceivedTickets(username, password, ticketOwner, ticketType, keyword);
```

#### دریافت تعداد تیکت های کاربران
```js
ticketSoapClient.GetReceivedTicketsCount(username, password, ticketType);
```

#### دریافت تیکت های ارسال شده
```js
ticketSoapClient.GetSentTickets(username, password, ticketOwner, ticketType, keyword);
```

#### دریافت تعداد تیکت های ارسال شده
```js
ticketSoapClient.GetSentTicketsCount(username, password, ticketType);
```


#### پاسخگویی به تیکت
```js
ticketSoapClient.ResponseTicket(username, password, ticketId, type, content, alertWithSms);
```
<div dir='rtl'>
  
### وب سرویس دفترچه تلفن

</div>

#### اضافه کردن گروه جدید
```js
contactsSoapClient.AddGroup(username, password, groupName, Descriptions, showToChilds);
```

#### اضافه کردن کاربر جدید
```js
contactsSoapClient.AddContact(username, password, options);

```

#### بررسی موجود بودن شماره در دفترچه تلفن
```js
contactsSoapClient.CheckMobileExistInContact(username, password, mobileNumber);
```

#### دریافت اطلاعات دفترچه تلفن
```js
contactsSoapClient.GetContacts(username, password, groupId, keyword, count);
```
#### دریافت گروه ها
```js
contactsSoapClient.GetGroups(username, password);
```
#### ویرایش مخاطب
```js
contactsSoapClient.ChangeContact(username, password, options);
```

#### حذف مخاطب
```js
contactsSoapClient.RemoveContact(username, password, mobilenumber);
```
#### دریافت اطلاعات مناسبت های فرد
```js
contactsSoapClient.GetContactEvents(username, password, contactId);
```

<div dir='rtl'>

### وب سرویس کاربران

</div>

#### ثبت فیش واریزی
```js
usersSoapClient.AddPayment(username, password, options);
```

#### اضافه کردن کاربر جدید در سامانه
```js
usersSoapClient.AddUser(username, password, options);

```

#### اضافه کردن کاربر جدید در سامانه(کامل)
```js
usersSoapClient.AddUserComplete(username, password, options);
```

#### اضافه کردن کاربر جدید در سامانه(WithLocation)
```js
usersSoapClient.AddUserWithLocation(username, password, options);
```
#### بدست آوردن ID کاربر
```js
usersSoapClient.AuthenticateUser(username, password);
```
#### تغییر اعتبار
```js
usersSoapClient.ChangeUserCredit(username, password, amount, description, targetUsername, GetTax);
```

#### فراموشی رمز عبور
```js
usersSoapClient.ForgotPassword(username, password, mobileNumber, emailAddress, targetUsername);
```
#### دریافت تعرفه پایه کاربر
```js
usersSoapClient.GetUserBasePrice(username, password, targetUsername);
```

#### دریافت اعتبار کاربر
```js
usersSoapClient.GetUserCredit(username, password, targetUsername);
```

#### دریافت مشخصات کاربر
```js
usersSoapClient.GetUserDetails(username, password, targetUsername);
```

#### دریافت شماره های کاربر
```js
usersSoapClient.GetUserNumbers(username, password);
```

#### دریافت تراکنش های کاربر
```js
usersSoapClient.GetUserTransactions(username, password, targetUsername, creditType, dateFrom, dateTo, keyword);
```

#### دریافت اطلاعات  کاربران
```js
usersSoapClient.GetUsers(username, password);
```


#### دریافت اطلاعات  فیلترینگ
```js
usersSoapClient.HasFilter(username, password, text);
```


####  حذف کاربر
```js
usersSoapClient.RemoveUser(username, password, targetUsername);
```


#### مشاهده استان ها
```js
usersSoapClient.GetProvinces(username, password);
```

#### مشاهده کد شهرستان 
```js
usersSoapClient.GetCities(username, password, provinceId);
```


#### مشاهده تاریخ انقضای کاربر 
```js
usersSoapClient.GetExpireDate(username, password);
```
