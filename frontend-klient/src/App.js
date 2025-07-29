import React, { useState } from 'react';
import './App.css';

function App() {
  const [selectedCategory, setSelectedCategory] = useState(null);
  const [cart, setCart] = useState([])
  const [showCart, setShowCart] = useState(false);
  
  const categories = [
    { id: 1, name: 'Wrapy', emoji: '🌯', description: 'Świeże wrapy z różnymi nadzieniami' },
    { id: 2, name: 'Burgery', emoji: '🍔', description: 'Socziste burgery na każdy gust' },
    { id: 3, name: 'Sałatki', emoji: '🥗', description: 'Zdrowe i świeże sałatki' }
  ];
  
  const menuItems = {
    1: [
      { id: 1, name: 'Wrap Klasyczny', price: 18, description: 'Kurczak, sałata, pomidor, ogórek, sos czosnkowy' },
      { id: 2, name: 'Wrap Wege', price: 16, description: 'Hummus, awokado, sałata, pomidor, ogórek, papryka' },
      { id: 3, name: 'Wrap Ostry', price: 19, description: 'Kurczak w ostrej marynacie, jalapeno, cebula, sos chipotle' }
    ],
    2: [
      { id: 4, name: 'Burger Klasyczny', price: 22, description: 'Wołowina, sałata, pomidor, cebula, sos burger' },
      { id: 5, name: 'Burger Wege', price: 20, description: 'Kotlet z quinoa, awokado, sałata, pomidor' }
    ],
    3: [
      { id: 6, name: 'Sałatka Cezar', price: 15, description: 'Sałata rzymska, kurczak, parmezan, grzanki' },
      { id: 7, name: 'Sałatka Grecka', price: 14, description: 'Pomidory, ogórki, oliwki, feta, czerwona cebula' }
    ]
  };

  const addToCart = (item) => {
    setCart(prevCart => {
      const existingItem = prevCart.find(cartItem => cartItem.id === item.id);

      if (existingItem) {
        return prevCart.map(cartItem =>
          cartItem.id === item.id
          ? { ...cartItem, quantity: cartItem.quantity + 1}
          : cartItem
        );
      } else {
        return [...prevCart, { ...item, quantity: 1 }];
      }
    });
  };
const removeFromCart = (itemId) => {
  setCart(prevCart => {
    return prevCart.map(cartItem => {
      if (cartItem.id === itemId) {
        if (cartItem.quantity > 1) {
          return { ...cartItem, quantity: cartItem.quantity - 1 };
        } else {
          return null;
        }
      }
      return cartItem;
    }).filter(item => item !== null);
  });
};
const clearCart = () => {
  setCart([]);
};


  return (
    <div className='App'>
      <header className='header'>
        <h1>Menu Restauracji</h1>
        <p>Wybierz danie</p>
        <div 
          className='cart-info'
          onClick={() => setShowCart(!showCart)}
        >
          Koszyk ({cart.reduce((total, item) => total + item.quantity, 0)})
        </div>
      </header>

      {showCart && (
        <div className='cart-overlay'>
          <div className='cart-modal'>
            <h3>Twój koszyk</h3>
            {cart.length === 0 ? (
              <p>Koszyk jest pusty</p>
            ) : (
              <>
                {cart.map(item => (
                  <div key={item.id} className='cart-item'>
                    <span>{item.name}</span>
                    <div className='quantity-controls'>
                    <button 
                      className='quantity-btn'
                      onClick={() => removeFromCart(item.id)}
                    >
                      -
                    </button>
                    <span>{item.quantity}x</span>
                    <button 
                      className='quantity-btn'
                      onClick={() => addToCart(item)}
                    >
                      +
                    </button>
                  </div>
                    <span>{item.price * item.quantity} zł</span>
                  </div>
                ))}
                <div className='cart-total'>
                  Suma: {cart.reduce((total, item) => total + (item.price * item.quantity), 0)} zł
                </div>
              </>
            )}
            <div className='cart-buttons'>
              <button 
                className='clear-cart-btn'
                onClick={clearCart}
                disabled={cart.length === 0}
              >
                Wyczyść koszyk
              </button>
              <button onClick={() => setShowCart(false)}>Zamknij</button>
            </div>
          </div>
        </div>
      )}


      <main className='main-content'>
        {!selectedCategory ? (
          <>
            <h2>Kategorie Menu</h2>
            <div className='categories'>
              {categories.map(category => (
                <div
                  key={category.id}
                  className='category-card'
                  onClick={() => setSelectedCategory(category.id)}
                >
                  <h3>{category.emoji} {category.name}</h3>
                  <p>{category.description}</p>
                </div>
            ))}
          </div>
          </>
        ) : (
          <>
            <button onClick={() => setSelectedCategory(null)}>
              Powrót do kategorii
            </button>
            <h2>{categories.find(c => c.id === selectedCategory)?.name}</h2>
            <div className='menu-items'>
              {menuItems[selectedCategory]?.map(item => (
                <div key={item.id} className='menu-item'>
                  <h3>{item.name}</h3>
                  <p>{item.description}</p>
                  <span className='price'>{item.price} zł</span>
                  <button
                    className='add-to-cart-btn'
                    onClick={() => addToCart(item)}
                  >
                    Dodaj do Koszyka
                  </button>
                </div>
              ))}
            </div>
          </>
        )}
      </main>
    </div>
  );
}

export default App;
