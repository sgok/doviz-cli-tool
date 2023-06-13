
# Döviz CLI Tool

![GitHub repo size](https://img.shields.io/github/repo-size/sgok/doviz-cli-tool) ![GitHub stars](https://img.shields.io/github/stars/sgok/doviz-cli-tool?style=social)

Döviz CLI Tool, anlık olarak döviz ve Bitcoin kurlarını takip etmenizi sağlayan bir komut satırı aracıdır. Program, Go programlama dili kullanılarak geliştirilmiştir ve standart olarak USD/TRY, EUR/TRY, GAU/TRY ve BTC/USD kurlarını sunar.

![Döviz CLI Tool Örneği](screenshot.png)

## Özellikler

-   Anlık olarak döviz kurlarını takip edebilme
-   USD/TRY, EUR/TRY, GAU/TRY ve BTC/USD kurlarını görüntüleme
-   Son güncellenen tarih ve saat bilgisini gösterme

## Gereksinimler

-   İşletim Sistemi: Windows, macOS, Linux
-   [Go](https://go.dev/dl/) sürüm 1.15 veya üstü

## Kurulum

1.  İlk olarak, Go'nun sisteminizde yüklü olduğundan emin olun. Eğer Go henüz yüklü değilse, aşağıdaki adımları takip ederek yükleyebilirsiniz:
    
    -   **Windows**:
        
        -   İndirme sayfasına gidin: [https://go.dev/dl/](https://go.dev/dl/)
        -   İndirilen `.msi` uzantılı dosyayı çalıştırın ve kurulum sihirbazını takip edin.
        -   Yükleme tamamlandıktan sonra, komut istemcisini yeniden başlatın ve `go version` komutunu çalıştırarak yüklü olan Go sürümünü kontrol edin.
    -   **macOS**:
        
        -   İndirme sayfasına gidin: [https://go.dev/dl/](https://go.dev/dl/)
        -   İndirilen `.pkg` uzantılı dosyayı çift tıklayarak kurulumu başlatın ve kurulum sihirbazını takip edin.
        -   Yükleme tamamlandıktan sonra, Terminal uygulamasını açın ve `go version` komutunu çalıştırarak yüklü olan Go sürümünü kontrol edin.
    -   **Linux**:
        
        -   Terminal uygulamasını açın ve aşağıdaki komutları sırasıyla çalıştırarak Go'yu yükleyin:
            
            
            `sudo apt update
            sudo apt install golang` 
            
        -   Yükleme tamamlandıktan sonra, terminalde `go version` komutunu çalıştırarak yüklü olan Go sürümünü kontrol edin.
2.  Döviz CLI Tool'u yerel bir dizine klonlayın:
    
    
    `git clone https://github.com/sgok/doviz-cli-tool.git` 
    
3.  Klonladığınız dizine gidin:
    
    
    `cd doviz-cli-tool` 
    
4.  Programı derleyin:
    
    
    `go build doviz.go` 
    

## Kullanım

1.  Döviz CLI Tool dizinindeyken, aşağıdaki komutu çalıştırarak programı başlatın:
    
    
    `./doviz` 
    
2.  Program, anlık döviz kurlarını ve Bitcoin değerini ekranda gösterecektir.
    

## Lisans

Bu proje MIT Lisansı ile lisanslanmıştır.

----------

**Not:** Bu projenin gerçek zamanlı döviz ve Bitcoin verilerine bağlı olduğunu unutmayın. Programın çıktısı, kullanılan API hizmetinin sağladığı verilere dayanmaktadır ve her zaman güncel olmayabilir. API adresini kendinize göre değiştirebilir veya kendi API adresinizi kullanabilirsiniz. API Endpoint'inin sonsuza kadar ayakta kalacağının garantisini veremeyiz.

Kendi API'sini kullanacaklar için örnek API çıktısı:
```json
{
    "time": {
        "updated": "Jun 13, 2023 21:31:03 UTC+03:00"
    },
    "disclaimer": "Bu veriler sadece bilgilendirme amaçlıdır.",
    "currencies": {
        "USD/TRY": {
            "name": "Dolar/Turkish Lira",
            "rate": "23,6197",
            "symbol": "₺"
        },
        "EUR/TRY": {
            "name": "Euro/Turkish Lira",
            "rate": "25,4847",
            "symbol": "₺"
        },
        "GAU/TRY": {
            "name": "Gold Gram/Turkish Lira",
            "rate": "1.475,04",
            "symbol": "₺"
        },
        "BTC/USD": {
            "name": "Bitcoin/Dolar",
            "rate": "25.898",
            "symbol": "$"
        }
    }
}
```