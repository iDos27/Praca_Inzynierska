import React, { useEffect, useState } from 'react';
import { apiService } from './api';
import './App.css';

function App() {
  const [selectedCategory, setSelectedCategory] = useState(null);
  const [cart, setCart] = useState([])
  const [showCart, setShowCart] = useState(false);
  const [showOrderForm, setShowOrderForm] = useState(false);
  const [orderData, setOrderData] = useState({
    orderType: 'table',
    tableNumber: ''
  });

  const [categories, setCategories] = useState([]);
  const [menuItems, setMenuItems] = useState({});
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  
  useEffect(() => {
    fetchCategories();
  }, []);

  const fetchCategories = async () => {
    try {
      setLoading(true);
      const response = await apiService.getCategories();
      setCategories(response.data);
    } catch (err) {
      setError('Błąd pobierania kategorii');
      console.error('Error fetching categories:', err);
    } finally {
      setLoading(false);
    }
  };
  const fetchMenuItems = async (categoryId) => {
    try {
      const response = await apiService.getMenuItems(categoryId);
      setMenuItems(prev => ({
        ...prev,
        [categoryId]: response.data
      }));
    } catch (err) {
      setError('Błąd pobierania menu');
      console.error('Error fetching menu items:', err);
    }
  };

  const handleCategorySelect = (categoryId) => {
    setSelectedCategory(categoryId);
    if (!menuItems[categoryId]) {
      fetchMenuItems(categoryId);
    }
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
  const handleOrderSubmit = async (e) => {
    e.preventDefault();

    if (orderData.orderType === 'table' && !orderData.tableNumber) {
      alert('Proszę podać numer stolika');
      return;
    }

    const orderPayload = {
      order_type: orderData.orderType,
      table_number: orderData.orderType === 'table' ? orderData.tableNumber : '',
      items: cart.map(item => ({
        id: item.id,
        quantity: item.quantity
      }))
    };

    try {
      const response = await apiService.createOrder(orderPayload);

      if (orderData.orderType === 'pickup') {
        alert(`Zamówienie przyjęte! Twój numer: ${response.data.order_number}. Poczekaj na wywołanie.`);
      } else {
        alert(`Zamówienie do stolika ${orderData.tableNumber} zostało przyjęte!`);
      }

      setCart([]);
      setOrderData({ orderType: 'table', tableNumber: '' });
      setShowOrderForm(false);
      setShowCart(false);
    } catch (err) {
      alert('Błąd przy składaniu zamówienia');
      console.error('Error creating order:', err);
    }
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
              <button 
                className='order-btn'
                onClick={() => setShowOrderForm(true)}
                disabled={cart.length === 0}
              >
                Złóż zamówienie
              </button>
              <button onClick={() => setShowCart(false)}>Zamknij</button>
            </div>
          </div>
        </div>
      )}

      {showOrderForm && (
        <div className='cart-overlay'>
          <div className='cart-modal'>
            <h3>Złóż zamówienie</h3>
            <form onSubmit={handleOrderSubmit}>
              <div className='form-group'>
                <label>Typ zamówienia *</label>
                <div className='radio-group'>
                  <label>
                    <input
                      type='radio'
                      name='orderType'
                      value='table'
                      checked={orderData.orderType === 'table'}
                      onChange={(e) => setOrderData({...orderData, orderType: e.target.value})}
                    />
                    Dostawa do stolika
                  </label>
                  <label>
                    <input
                      type='radio'
                      name='orderType'
                      value='pickup'
                      checked={orderData.orderType === 'pickup'}
                      onChange={(e) => setOrderData({...orderData, orderType: e.target.value, tableNumber: ''})}
                    />
                    Odbiór przy ladzie
                  </label>
                </div>
              </div>

              {orderData.orderType === 'table' && (
                <div className='form-group'>
                  <label>Numer stolika *</label>
                  <input
                    type='number'
                    value={orderData.tableNumber}
                    onChange={(e) => setOrderData({...orderData, tableNumber: e.target.value})}
                    placeholder='np. 12'
                    min='1'
                    required
                  />
                </div>
              )}
              
              <div className='order-summary'>
                <p>Wartość zamówienia: <strong>{cart.reduce((total, item) => total + (item.price * item.quantity), 0)} zł</strong></p>
                <p>Liczba pozycji: <strong>{cart.reduce((total, item) => total + item.quantity, 0)}</strong></p>
              </div>
              
              <div className='form-buttons'>
                <button type='button' onClick={() => setShowOrderForm(false)}>Anuluj</button>
                <button type='submit' className='submit-btn'>Potwierdź zamówienie</button>
              </div>
            </form>
          </div>
        </div>
      )}
      
      {loading && <div className="loading">Ładowanie...</div>}
      {error && <div className="error">{error}</div>}

      <main className='main-content'>
        {!selectedCategory ? (
          <>
            <h2>Kategorie Menu</h2>
            <div className='categories'>
              {categories.map(category => (
                <div
                  key={category.id}
                  className='category-card'
                  onClick={() => handleCategorySelect(category.id)}
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
