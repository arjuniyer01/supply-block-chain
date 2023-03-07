import streamlit as st  
from streamlit_qrcode_scanner import qrcode_scanner  

qr_code = qrcode_scanner(key="qr_code_scanner")  

if qr_code:  
  st.write(qr_code) 