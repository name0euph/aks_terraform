import { useEffect } from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { Auth } from './components/Auth';
import { Todo } from './components/Todo';
import axios from 'axios';
import { CsrfToken } from './types';

function App() {
  useEffect(() => {
    axios.defaults.withCredentials = true;

    const getCsrfToken = async () => {
      // CSRF トークンを取得するエンドポイントにアクセス
      const { data } = await axios.get<CsrfToken>(
        `${process.env.REACT_APP_API_URL}/csrf`
      );
      // CSRF トークンをリクエストヘッダーに設定
      axios.defaults.headers.common['X-CSRF-Token'] = data.csrf_token;
    } 
    getCsrfToken();
  }, []);

  return (
    <Router>
      <Routes>
        <Route path="/" element={<Auth />} />
        <Route path="/todo" element={<Todo />} />
      </Routes>
    </Router>
  );
}

export default App;
