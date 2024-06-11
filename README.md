Program, kullanıcıya bir menü sunarak stack veya queue işlemleri arasında seçim yapmasını sağlar. Kullanıcı, stack veya queue'ya eleman ekleme, belirli bir elemanı bulma ve silme, en önündeki elemanı görüntüleme ve tüm elemanları listeleme gibi işlemler yapabilir.
Ana Menü

Ana menü, programın başlangıç noktasıdır ve kullanıcıya üç seçenek sunar: Stack işlemleri, Queue işlemleri ve programdan çıkış. Kullanıcı, ilgili seçeneği girerek ilgili menüye yönlendirilir veya programdan çıkış yapar.
Stack İşlemleri Menüsü

Stack işlemleri menüsü, kullanıcıya beş seçenek sunar:

    Stack'e eleman ekleme
    Stack'te bir eleman bulup silme
    Stack'te bir eleman bulup gösterme
    Stack'teki tüm elemanları listeleme
    Ana menüye dönme

Stack'e Eleman Ekleme

Bu işlev, kullanıcıdan bir değer alır ve bu değeri stack'e ekler. Stack, son girilen değerin ilk çıkarıldığı (LIFO - Last In, First Out) bir veri yapısıdır.
Stack'te Bir Eleman Bulup Silme

Bu işlev, kullanıcıdan bir değer alır, bu değeri stack'te arar ve bulursa siler. Eğer değer bulunamazsa, kullanıcıya değer bulunamadığı bildirilir.
Stack'te Bir Eleman Bulup Gösterme

Bu işlev, kullanıcıdan bir değer alır ve bu değerin stack'te olup olmadığını kontrol eder. Eğer değer bulunursa, değerin bulunduğu bildirilir; aksi takdirde, değerin bulunamadığı bildirilir.
Stack'teki Tüm Elemanları Listeleme

Bu işlev, stack'teki tüm elemanları listeler. Eğer stack boş ise, kullanıcıya stack'in boş olduğu bildirilir.
Queue İşlemleri Menüsü

Queue işlemleri menüsü, kullanıcıya beş seçenek sunar:

    Queue'ya eleman ekleme
    Queue'dan eleman kaldırma
    Queue'nun en önündeki elemanı gösterme
    Queue'daki tüm elemanları listeleme
    Ana menüye dönme

Queue'ya Eleman Ekleme

Bu işlev, kullanıcıdan bir değer alır ve bu değeri queue'ya ekler. Queue, ilk girilen değerin ilk çıkarıldığı (FIFO - First In, First Out) bir veri yapısıdır.
Queue'dan Eleman Kaldırma

Bu işlev, queue'nun en önündeki elemanı kaldırır ve kullanıcıya bildirir. Eğer queue boş ise, kullanıcıya queue'nun boş olduğu bildirilir.
Queue'nun En Önündeki Elemanı Gösterme

Bu işlev, queue'nun en önündeki elemanı gösterir. Eğer queue boş ise, kullanıcıya queue'nun boş olduğu bildirilir.
Queue'daki Tüm Elemanları Listeleme

Bu işlev, queue'daki tüm elemanları listeler. Eğer queue boş ise, kullanıcıya queue'nun boş olduğu bildirilir.
Yardımcı Fonksiyonlar
Kullanıcıdan Girdi Alma

Kullanıcıdan girdi almak için kullanılan yardımcı bir işlev, kullanıcıya bir prompt gösterir ve girdiyi alarak döndürür. Girdinin başındaki ve sonundaki boşluklar temizlenir.
Sonuç

Bu uygulama, kullanıcıların Stack ve Queue veri yapılarını kullanarak çeşitli işlemleri gerçekleştirmelerine olanak tanır. Stack ve Queue, bilgisayar bilimlerinde yaygın olarak kullanılan veri yapılarıdır ve bu uygulama, bu yapıların temel işlemlerini gerçekleştirmek için kullanıcı dostu bir arayüz sağlar. Program, kullanıcı girişlerini alarak ilgili işlemleri gerçekleştirir ve sonuçları kullanıcıya bildirir. Bu tür uygulamalar, veri yapılarını anlamak ve kullanmak için temel bir eğitim aracı olarak değerlendirilebilir.

![Ekran görüntüsü 2024-06-11 194439](https://github.com/arazumut/learnGoandStack-Qeue.go/assets/150933483/3e683782-ae34-44a4-8407-7968fa2e6541)
![Ekran görüntüsü 2024-06-11 194331](https://github.com/arazumut/learnGoandStack-Qeue.go/assets/150933483/cec7d8ea-10ac-4288-a9db-b7215deaac8b)
![Ekran görüntüsü 2024-06-11 194307](https://github.com/arazumut/learnGoandStack-Qeue.go/assets/150933483/b4f43b3f-409d-4137-a23b-3cdfdefccdfb)
![Ekran görüntüsü 2024-06-11 194250](https://github.com/arazumut/learnGoandStack-Qeue.go/assets/150933483/249048cb-2bed-45b0-83bf-43c467e60726)
![Ekran görüntüsü 2024-06-11 194239](https://github.com/arazumut/learnGoandStack-Qeue.go/assets/150933483/f8f15c39-9862-40bf-8a6f-a345f0db85bd)
