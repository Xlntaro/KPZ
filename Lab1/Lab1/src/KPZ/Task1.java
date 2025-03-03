package KPZ;
import java.util.ArrayList;
import java.util.List;

// Клас Money для зберігання ціни у форматі "гривні + копійки"
class Money {
    private int wholePart;
    private int cents;

    public Money(int wholePart, int cents) {
        this.wholePart = wholePart;
        this.cents = cents;
    }

    public int getWholePart() {
        return wholePart;
    }

    public int getCents() {
        return cents;
    }

    public void setMoney(int wholePart, int cents) {
        this.wholePart = wholePart;
        this.cents = cents;
    }

    public void display() {
        System.out.printf("%d.%02d UAH%n", wholePart, cents);
    }
}

// Клас Product для збереження товарів
class Product {
    private String name;
    private Money price;

    public Product(String name, Money price) {
        this.name = name;
        this.price = price;
    }

    public void reducePrice(int amount) {
        int newCents = price.getCents() - amount;
        if (newCents < 0) {
            price.setMoney(price.getWholePart() - 1, 100 + newCents);
        } else {
            price.setMoney(price.getWholePart(), newCents);
        }
    }

    public void displayProduct() {
        System.out.print("Product: " + name + ", Price: ");
        price.display();
    }
}

// Клас Warehouse для збереження списку товарів
class Warehouse {
    private List<Product> products = new ArrayList<>();

    public void addProduct(Product product) {
        products.add(product);
    }

    public void displayProducts() {
        for (Product product : products) {
            product.displayProduct();
        }
    }
}

// Клас Reporting для генерації звітів
class Reporting {
    public static void generateInventoryReport(Warehouse warehouse) {
        System.out.println("=== Inventory Report ===");
        warehouse.displayProducts();
    }
}
