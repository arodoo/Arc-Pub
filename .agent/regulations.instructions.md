# GDD: MECÁNICAS DE ECONOMÍA Y BALANCE DE JUEGO

## 1. NÚCLEO DE JUEGO: GUERRA DE RUTAS COMERCIALES
**Concepto Fundamental:** Riesgo vs. Recompensa Geográfico.
El comercio no es instantáneo (teletransportado). Los recursos de alto valor (Tier 2 y 3) se generan exclusivamente en zonas de conflicto (Mapas PvP o Fronterizos).

* **Mecánica de Transporte Físico:** Para vender estos recursos en la Casa de Subastas (ubicada en Zonas Seguras/Capitales), el jugador debe transportarlos en su inventario a través del mapa.
* **Factor de Riesgo:** Si el jugador es abatido en el trayecto, un porcentaje significativo de la carga (ej. 30-50%) cae al suelo y puede ser robado por el atacante.
* **Consecuencia Social:** Esto incentiva la creación de **Escoltas y Caravanas**. Los jugadores solitarios correrán riesgos altos, mientras que los clanes cobrarán por protección. Genera contenido orgánico sin necesidad de misiones programadas.

## 2. FILTRO DE MERCADO: LA "LICENCIA DE COMERCIO" (Token)
**Concepto:** Economía de Desagüe (Sink Economy).
Para evitar la hiperinflación y el spam de objetos basura en la subasta, introducimos una barrera de entrada al comercio.

* **El Token:** Un objeto llamado "Licencia de Comercio".
* **La Regla:** Publicar una subasta o iniciar un intercambio directo (P2P) consume 1 Licencia.
* **Obtención Controlada:**
    1.  **Drop Raro:** Los Jefes de Mundo o Mobs de Élite tienen una probabilidad baja (0.5% - 1%) de soltarlo. Fomenta el juego activo (PvE).
    2.  **Microtransacción Barata:** Se puede adquirir por un precio muy bajo (ej. $0.50 USD).
* **Impacto Económico:**
    * **Anti-Bots:** Crear miles de subastas basura deja de ser rentable para los bots, ya que cada publicación tiene un costo real o de tiempo.
    * **Valorización:** Los jugadores solo subastarán objetos que valgan más que el costo de la Licencia. Limpia el mercado de "basura".

## 3. ANTI-PAY-TO-WIN: VINCULACIÓN DE OBJETOS (Soulbound)
**Concepto:** Mérito sobre Capacidad de Compra.
Evitamos que un jugador rico compre el mejor equipamiento del juego el día 1 sin haber jugado.

* **Clasificación de Ítems:**
    * **Tier C (Común):** Libre comercio.
    * **Tier B (Raro):** Se vincula al equipar (BoE). Una vez te lo pones, nunca más puedes venderlo.
    * **Tier A/S (Épico/Legendario):** Se vincula al recoger (BoP). Solo quien mata al jefe puede usar el objeto.
* **Consecuencia:** Ver a un jugador con una armadura dorada significa que ese jugador (o su equipo) derrotó al dragón, no que usó su tarjeta de crédito. Preserva el prestigio visual.

## 4. SUBASTA CONVENCIONAL CON CIERRE DINÁMICO (Soft Closing)
**Concepto:** Justicia Temporal.
Resolvemos el problema del "Sniping" (bots robando subastas en el último segundo) manteniendo el formato clásico de pujas visibles.

* **Mecánica Base:** Subasta pública, precio visible, gana el mejor postor.
* **Regla Anti-Sniping:** Si se realiza una puja cuando quedan menos de **X minutos** (ej. 2 min) para finalizar, el reloj se reinicia automáticamente a **X minutos**.
* **Ejemplo Práctico:**
    * La subasta termina a las 12:00.
    * Un bot puja a las 11:59:59 para robar el ítem.
    * El sistema detecta la puja tardía y extiende el cierre a las 12:02:00.
    * El jugador humano recibe una notificación y tiene 2 minutos para contraofertar.
* **Resultado:** La subasta solo termina cuando nadie está dispuesto a pagar más, eliminando la ventaja técnica de los scripts automatizados.

## 5. CONTROL DE TRANSFERENCIA: EL IMPUESTO "ROBIN HOOD"
**Concepto:** Fricción en el Mercado Negro.
Dificultamos la venta ilegal de oro (RMT) y el "power-leveling" financiero desmedido.

* **Detección:** El sistema analiza la diferencia de **Valor Patrimonial (Net Worth)** o Nivel entre dos jugadores que comercian.
* **Impuesto Progresivo:**
    * Jugador Nv. 50 comercia con Jugador Nv. 50 -> Impuesto estándar (5%).
    * Jugador Nv. 50 regala 1 millón a Jugador Nv. 1 -> Impuesto de Castigo (90%).
* **Narrativa:** "La burocracia de la Federación cobra aranceles por donaciones sospechosas".
* **Efecto:** El vendedor de oro ilegal tiene que farmear 10 millones para entregar 1 millón al cliente, destruyendo su margen de beneficio.

## 6. SISTEMA DE JUSTICIA PVP: ESTADO CRIMINAL
**Concepto:** Consecuencias, no prohibiciones.
El PvP está permitido, pero el abuso tiene un coste sistémico severo.

* **El abuso:** Matar jugadores de "Rango Gris" (ej. 10 niveles inferiores al tuyo) o en zonas no designadas.
* **La Marca:** El agresor recibe el estado "Criminal" por un tiempo (ej. 4 horas de juego activo).
* **Castigos del Estado Criminal:**
    1.  **Exilio Económico:** Los NPC de las estaciones seguras (Vendedores, Reparadores) se niegan a comerciar con criminales.
    2.  **Caza Abierta:** El criminal aparece marcado en el mapa global. Matarlo otorga recompensa y NO genera estado criminal al justiciero.
    3.  **Pérdida de Equipo:** Si un criminal muere, tiene una probabilidad mayor de soltar objetos equipados (Drop rate aumentado).