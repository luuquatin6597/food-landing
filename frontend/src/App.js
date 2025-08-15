import React, { useState, useEffect } from 'react';
import axios from 'axios';
import './App.css';

// URL của Backend API, bạn sẽ thay thế bằng IP của VM sau này
const API_URL = 'http://YOUR_GCP_VM_IP:8080/api/foods';

function FoodCard({ food }) {
  return (
    <div className="food-card" style={{ borderTop: `5px solid ${food.color}` }}>
      <img src={food.image_url} alt={food.name} className="food-image" />
      <div className="food-content">
        <h3>{food.name} - <span className="food-region">{food.region}</span></h3>
        <p className="food-description">{food.description}</p>
        <p><strong>Nguyên liệu:</strong> {food.ingredients}</p>
        <p className="food-price">Giá: {food.price}</p>
      </div>
    </div>
  );
}

function App() {
  const [foods, setFoods] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchFoods = async () => {
      try {
        const response = await axios.get(API_URL);
        setFoods(response.data);
      } catch (err) {
        setError('Không thể tải dữ liệu món ăn. Vui lòng kiểm tra lại kết nối API.');
        console.error(err);
      } finally {
        setLoading(false);
      }
    };

    fetchFoods();
  }, []);

  return (
    <div className="App">
      <header className="App-header">
        <h1>Khám Phá Ẩm Thực Việt Nam 🍜</h1>
      </header>
      <main className="food-grid">
        {loading && <p>Đang tải danh sách món ăn...</p>}
        {error && <p className="error-message">{error}</p>}
        {foods.map(food => (
          <FoodCard key={food.id} food={food} />
        ))}
      </main>
    </div>
  );
}

export default App;