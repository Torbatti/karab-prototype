# Karab Prototype (Htmx + Go)
- 1402/05/19-1402/05/21

- یک ساید پروژه برای نشون دادن توانایی زبان گو و لایبرری اچ تی ام اکس 

- اچ تی ام اکس این قابلیت رو ایجاد کرده که به جای دریافت فایل جی سان از بک اند و تبدیل کردنش به دام , یک فایل اچ تی امل از سرور دریافت کنید و اون رو جایی که مشخص میکنید تزریق کنید

- برای دیدن کارکرد بک اند و فرانات اند  میتونید ویدیو زیر رو ببینید
[ویدیوی شو کیس](/Screen/ScreenCast.webm)

![ویدیوی شو کیس](/Screen/ScreenCast.gif)

# Design | طراحی
- مرحله طراحی که مهم ترین بخش ساخت هر برنامه ایه رو با نرم افزار penpot انجام دادم 
- میتونید فایل پن پات رو دانلود کنید و ۳ صفحه  طراحی شده رو ببینید برای آسانی کار اسکرین شات هاشو براتون این پایین گذاشتم
[Karab Penpot File](/penpot/Karab.penpot)

/index
![](/Screen/PenPot1.png)

/jobOpportunities
![](/Screen/PenPot2.png)

/search
![](/Screen/PenPot3.png)
# Implementation | پیاده سازی
/index.html
![](/Screen/Html1.png)

/search.html
![](/Screen/Html2.png)
# Backend Terminal Logs | لاگ بکند
- برای اینکه بدونی کی و کجا سرور کرش شده یه لاگینگ ساده زدم که اماده مرحله پروداکشن اصلا نیست ولی کار راه اندازه
![](/Screen/BackendLogs.png)


# نحوه ی استفاده
پیش نیاز: زبان گو و گیت
```
git clone https://github.com/Torbatti/karab-prototype.git
cd karab-prototype
go run ./backend
```
# نکته های قابل توجه
- برای پروداکشن اماده نیست ! 
- متصل به دیتا بیس نسیت و داده ها هارد کود شده هستن
