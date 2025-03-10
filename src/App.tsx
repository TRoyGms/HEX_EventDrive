import './App.css';
import axios from 'axios';
import { useEffect, useState } from 'react';
import Swal from 'sweetalert2';

function App() {
  const [socket, setSocket] = useState<WebSocket | null>(null);
  const [isLoading, setIsLoading] = useState(false);

  const api1Url = import.meta.env.VITE_API1_URL;      
  const api2SocketUrl = import.meta.env.VITE_HOST_SOCKET; 

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setIsLoading(true);
  
    const formData = new FormData(e.currentTarget);
  
    try {
      const data = {
        name: formData.get("name"),
        amount: Number(formData.get("amount")),
        description: formData.get("description") || "", 
      };

      await axios.post(`${api1Url}`, data);
      console.log("Pago enviado a API1 (OrderService)");
    } catch (error) {
      console.error(error);
      Swal.fire({
        icon: "error",
        title: "Error",
        text: "Hubo un problema al procesar tu pago.",
      });
    } finally {
      setIsLoading(false);
    }
  };
  
  useEffect(() => {
    if (!socket) {
      const wsUrl = api2SocketUrl.endsWith('/') ? `${api2SocketUrl}ws` : `${api2SocketUrl}/ws`;
      console.log("Intentando conectar al WebSocket en:", wsUrl);

      const ws = new WebSocket(wsUrl);

      ws.onopen = () => {
        console.log("Conectado al WebSocket de PaymentService");
      };

      ws.onerror = (error) => {
        console.error("Error en WebSocket:", error);
      };

      ws.onmessage = (event) => {
        console.log("Mensaje recibido desde WebSocket:", event.data);
        
        try {
          // Convertir el mensaje de JSON a objeto
          const paymentData = JSON.parse(event.data);

          // Mostrar SweetAlert con los datos del pago
          Swal.fire({
            icon: "success",
            title: "¡Pago confirmado!",
            text: `Gracias ${paymentData.name} por tu pago de $${paymentData.amount}`,
            confirmButtonText: "OK",
          });

        } catch (error) {
          console.error("Error procesando mensaje WebSocket:", error);
        }
      };

      ws.onclose = () => {
        console.log("WebSocket cerrado, intentando reconectar en 3s...");
        setTimeout(() => setSocket(null), 3000);
      };

      setSocket(ws);
    }
  }, [socket]);

  return (
    <>
      <div className='main'>
        <h1>Pago</h1><br />
        <form onSubmit={handleSubmit}>
          <label htmlFor="nombre">
            <p>Nombre:</p>
            <input id='nombre' type="text" name='name' required />
          </label><br /><br />
          <label htmlFor="monto">
            <p>Monto:</p>
            <input id='monto' type="number" name='amount' required />
          </label><br /><br />
          <label htmlFor="descripcion">
            <p>Descripción (opcional):</p>
            <input id="descripcion" type="text" name="description" />
          </label><br /><br />
          <button disabled={isLoading}>
            {isLoading ? "Procesando..." : "PAGAR"}
          </button>
        </form>
      </div>
    </>
  );
}

export default App;