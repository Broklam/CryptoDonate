import logo from './logo.svg';
import './App.css';
import NavBar from './components/NavBar';
import Present from './components/Present';
import Advice from './components/Advice';
import Footer from './components/Footer';
import { BrowserRouter, Route, Switch } from 'react-router-dom';
import FAQ from './pages/FAQ';


function App() {
  return (
    <div className='wrapper'>
      <BrowserRouter>
      <Switch>
      <Route path="/faq">
        <FAQ />
        </Route>
      <NavBar />
      <Present />
      <Advice />
      <Footer />
      </Switch>
      </BrowserRouter>
    </div>
  );
}

export default App;
