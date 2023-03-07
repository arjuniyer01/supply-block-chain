import streamlit as st  
from streamlit_qrcode_scanner import qrcode_scanner
from urllib.parse import urlparse
from urllib.parse import parse_qs

genesis_id = st.experimental_get_query_params()["id"]

def get_blockchain(genesis_id):
  # TODO: Use Go
  return

def get_blocks(block_hash):
  # TODO: Use Go
  return

st.title("Welcome to supply-block-chain :chains: :truck:")
st.markdown("### It seems you are in physical possession of a product. Please scan the QR code again to get the blockchain details for this product.")

st.title("Scan QR Code")
qr_code = qrcode_scanner(key="qr_code_scanner")  

if qr_code:
  captured_value = parse_qs(urlparse(qr_code).query)['id'][0]
  if genesis_id == captured_value:
    st.success("QR code scanned successfully!")
    selection = st.selectbox("Actions", ["View Blockchain", "I am the next owner", "Report product as lost"])
    if st.button("Process Action"):
      if selection == "View Blockchain":
        st.success("Viewing blockchain")
        get_blockchain(genesis_id)
      elif selection == "I am the next owner":
        st.success("I am the next owner")
        get_blocks(genesis_id)
      elif selection == "Report product as lost":
        st.success("Report product as lost")
        get_blocks(genesis_id)