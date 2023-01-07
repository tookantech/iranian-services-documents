# melipayamak Go (golang)

<div dir='rtl'>

### معرفی وب سرویس ملی پیامک
ملی پیامک یک وب سرویس کامل برای ارسال و دریافت پیامک و پیامک صوتی و مدیریت کامل خدمات دیگر است که براحتی میتوانید از آن استفاده کنید.

<hr>

### نصب

<p>قبل از نصب نیاز به ثبت نام در سایت ملی پیامک دارید.</p>

[**ثبت نام به همراه دریافت 200 پیامک هدیه جهت تست وبسرویس**](http://www.melipayamak.com/)

</div>

<div dir='rtl'>
  
#### نحوه استفاده

نمونه کد برای ارسال پیامک

</div>


```go
username := "username";
password := "password";
from := "5000...";
to := "09123456789";
text := "تست وب سرویس ملی پیامک";
isFlash := false;
SendSimpleSMS2(username, password, to, from, text, isFlash);
//یا برای ارسال به مجموعه ای از مخاطبین
SendSimpleSMS(username, password, to []string, from, text, isFlash);
```

<div dir='rtl'>

از آنجا که وب سرویس ملی پیامک تنها محدود به ارسال پیامک نیست شما از طریق متدهای زیر میتوانید به وب سرویس ها دسترسی کامل داشته باشید

#### تفاوت های وب سرویس پیامک rest و soap

از آنجا که ملی پیامک وب سرویس کاملی رو در اختیار توسعه دهندگان میگزارد برای راحتی کار با وب سرویس پیامک علاوه بر وب سرویس اصلی soap وب سرویس rest رو هم در اختیار توسعه دهندگان گزاشته شده تا راحتتر بتوانند با وب سرویس کار کنند. تفاوت اصلی این دو در تعداد متد هاییست که میتوانید با آن کار کنید. برای کار های پایه میتوان از وب سرویس rest استفاده کرد برای دسترسی بیشتر و استفاده پیشرفته تر نیز باید از وب سرویس باید از وب سرویس soap استفاده کرد. جهت مطالعه بیشتر وب سرویس ها به قسمت وب سرویس پنل خود مراجعه کنید.

<hr/>


###  اطلاعات بیشتر

برای مطالعه بیشتر و دریافت راهنمای وب سرویس ها و آشنایی با پارامتر های ورودی و خروجی وب سرویس به صفحه معرفی
[وب سرویس ملی پیامک](https://github.com/Melipayamak/Webservices)
مراجعه نمایید .


<hr/>

</div>


<div dir='rtl'>

### وب سرویس پیامک

متدهای وب سرویس:

</div>

#### ارسال

```go
Send(username, password, to, from, text, isFlash);
SendSimpleSMS(username, password, to []string, from, text, isFlash);
```

#### ارسال از خط خدماتی اشتراکی

```go
SendByBaseNumber(username, password, text, to, bodyId);
//soap
SendByBaseNumber2(username, password, text, to, bodyId);
```

#### دریافت وضعیت ارسال
```go
GetDelivery(username, password, recId);
GetDeliveries(username, password, recIds []string);
```

#### لیست پیامک ها

```go
GetMessages(username, password, location, index, count, from);
getMessages(username, password, location, from, index, count);
// جهت دریافت به صورت رشته ای
GetMessagesByDate(username, password, location, from, index, count, dateFrom, dateTo);
//جهت دریافت بر اساس تاریخ
GetUsersMessagesByDate(username, password, location, from, index, count, dateFrom, dateTo);
// جهت دریافت پیام های کاربران بر اساس تاریخ 
```

#### موجودی
```go
GetCredit(username, password);
```

#### تعرفه پایه / دریافت قیمت قبل از ارسال
```go
GetBasePrice(username, password);
GetSmsPrice(username, password, irancellCount, mtnCount, from, text);
```
#### لیست شماره اختصاصی
```go
GetUserNumbers(username, password);
```

#### بررسی تعداد پیامک های دریافتی
```go
GetInboxCount(username, password, isRead);
//پیش فرض خوانده نشده 
```

#### ارسال پیامک پیشرفته
```go
SendSms(username, password, to []string, from, text, isflash, udh, recId []string, status []string);
```

#### مشاهده مشخصات پیام
```go
GetMessagesReceptions(username, password, msgId, fromRows);
```


#### حذف پیام دریافتی
```go
RemoveMessages2(username, password, location, msgIds);
```


#### ارسال زماندار
```go
AddSchedule(username, password, to, from, text, isflash, scheduleDateTime, period);
```

#### ارسال زماندار متناظر
```go
AddMultipleSchedule(username, password, to []string, from, text []string, isflash, scheduleDateTime []string, period);
```


#### ارسال سررسید
```go
AddUsance(username, password, to, from, text, isflash, scheduleStartDateTime, repeatAfterDays, scheduleEndDateTime);
```

#### مشاهده وضعیت ارسال زماندار
```go
GetScheduleStatus(username, password, schId);
```

#### حذف پیامک زماندار
```go
RemoveSchedule(username, password, schId);
```

<div dir='rtl'>
  
### وب سرویس پیامک صوتی

</div>

####  ارسال پیامک همراه با تماس صوتی
```go
SendSMSWithSpeechText(username, password, smsBody, speechBody, from, to);
```

####  ارسال پیامک همراه با تماس صوتی به صورت زمانبندی
```go
SendSMSWithSpeechTextBySchduleDate(username, password, smsBody, speechBody, from, to, scheduleDate);
```

####  دریافت وضعیت پیامک همراه با تماس صوتی 
```go
GetSendSMSWithSpeechTextStatus(username, password, recId);
```

#### تماس انبوه زماندار
```go
SendBulkSpeechText(username, password, title, body, receivers, DateToSend, repeatCount);
```


#### تماس انبوه زماندار با انتخاب فایل
```go
SendBulkVoiceSMS(username, password, title, voiceFileId, receivers, DateToSend, repeatCount);
```


#### آپلود فایل صوتی
```go
UploadVoiceFile(username, password, title, file);
```

<div dir='rtl'>
  
### وب سرویس ارسال انبوه/منطقه ای

</div>

#### دریافت شناسه شاخه های بانک شماره
```go
GetBranchs(username, password, owner);
```


#### اضافه کردن یک بانک شماره جدید
```go
AddBranch(username, password, branchName, owner);
```

#### اضافه کردن شماره به بانک
```go
AddNumber(username, password, branchId, mobileNumbers []string);
```

#### حذف یک بانک
```go
RemoveBranch(username, password, branchId);
```

#### ارسال انبوه از طریق بانک
```go
AddBulk(username, password, from, branch, bulkType, title, message, rangeFrom, rangeTo, DateToSend, requestCount, rowFrom);
```

#### تعداد شماره های موجود
```go
GetBulkCount(username, password, branch, rangeFrom, rangeTo);
```

#### گزارش گیری از ارسال انبوه
```go
GetBulkReceptions(username, password, bulkId, fromRows);
```


#### تعیین وضعیت ارسال 
```go
GetBulkStatus(username, password, bulkId, sent, failed, status);
```

#### تعداد ارسال های امروز
```go
GetTodaySent(username, password);
```

#### تعداد ارسال های کل

```go
GetTotalSent(username, password);
```

#### حذف ارسال منطقه ای
```go
RemoveBulk(username, password, bulkId);
```

#### ارسال متناظر
```go
SendMultipleSMS(username, password, to []string, from, text, isflash, udh, recId []string, status);
```

#### نمایش دهنده وضعیت گزارش گیری

```go
UpdateBulkDelivery(username, password, bulkId);
```
<div dir='rtl'>
  
### وب سرویس تیکت

</div>

#### ثبت تیکت جدید
```go
AddTicket(username, password, title, content, aletWithSms);
```

#### جستجو و دریافت تیکت ها

```go
GetReceivedTickets(username, password, ticketOwner, ticketType, keyword);
```

#### دریافت تعداد تیکت های کاربران
```go
GetReceivedTicketsCount(username, password, ticketType);
```

#### دریافت تیکت های ارسال شده
```go
GetSentTickets(username, password, ticketOwner, ticketType, keyword);
```

#### دریافت تعداد تیکت های ارسال شده
```go
GetSentTicketsCount(username, password, ticketType);
```


#### پاسخگویی به تیکت
```go
ResponseTicket(username, password, ticketId, type, content, alertWithSms);
```

<div dir='rtl'>
  
### وب سرویس دفترچه تلفن

</div>

#### اضافه کردن گروه جدید
```go
AddGroup(username, password, groupName, Descriptions, showToChilds);
```

#### اضافه کردن کاربر جدید
```go
AddContact(username, password, options);

```

#### بررسی موجود بودن شماره در دفترچه تلفن
```go
CheckMobileExistInContact(username, password, mobileNumber);
```

#### دریافت اطلاعات دفترچه تلفن
```go
GetContacts(username, password, groupId, keyword, count);
```
#### دریافت گروه ها
```go
GetGroups(username, password);
```
#### ویرایش مخاطب
```go
ChangeContact(username, password, options);
```

#### حذف مخاطب
```go
RemoveContact(username, password, mobilenumber);
```
#### دریافت اطلاعات مناسبت های فرد
```go
GetContactEvents(username, password, contactId);
```

<div dir='rtl'>

### وب سرویس کاربران

</div>

#### ثبت فیش واریزی
```go
AddPayment(username, password, options);
```

#### اضافه کردن کاربر جدید در سامانه
```go
AddUser(username, password, options);

```

#### اضافه کردن کاربر جدید در سامانه(کامل)
```go
AddUserComplete(username, password, options);
```

#### اضافه کردن کاربر جدید در سامانه(WithLocation)
```go
AddUserWithLocation(username, password, options);
```
#### بدست آوردن ID کاربر
```go
AuthenticateUser(username, password);
```
#### تغییر اعتبار
```go
ChangeUserCredit(username, password, amount, description, targetUsername, GetTax);
```

#### فراموشی رمز عبور
```go
ForgotPassword(username, password, mobileNumber, emailAddress, targetUsername);
```
#### دریافت تعرفه پایه کاربر
```go
GetUserBasePrice(username, password, targetUsername);
```

#### دریافت اعتبار کاربر
```go
GetUserCredit(username, password, targetUsername);
```

#### دریافت مشخصات کاربر
```go
GetUserDetails(username, password, targetUsername);
```

#### دریافت شماره های کاربر
```go
GetUserNumbers(username, password);
```

#### دریافت تراکنش های کاربر
```go
GetUserTransactions(username, password, targetUsername, creditType, dateFrom, dateTo, keyword);
```

#### دریافت اطلاعات  کاربران
```go
GetUsers(username, password);
```


#### دریافت اطلاعات  فیلترینگ
```go
HasFilter(username, password, text);
```


####  حذف کاربر
```go
RemoveUser(username, password, targetUsername);
```


#### مشاهده استان ها
```go
GetProvinces(username, password);
```

#### مشاهده کد شهرستان 
```go
GetCities(username, password, provinceId);
```


#### مشاهده تاریخ انقضای کاربر 
```go
GetExpireDate(username, password);
```
