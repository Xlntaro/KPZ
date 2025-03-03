package KPZ;

public class Task2 {
    public static void main(String[] args) {
        // Створення продуктів
        Money price1 = new Money(10, 50);
        Product product1 = new Product("Laptop", price1);

        Money price2 = new Money(5, 75);
        Product product2 = new Product("Mouse", price2);

        Money price3 = new Money(2, 30);
        Product product3 = new Product("Keyboard", price3);

        // Створення складу
        Warehouse warehouse = new Warehouse();
        warehouse.addProduct(product1);
        warehouse.addProduct(product2);
        warehouse.addProduct(product3);

        // Генерація звіту
        Reporting.generateInventoryReport(warehouse);
    }
}

