git clone https://github.com/gpBlockchain/ckb-rpc-mock-data.git
cd ckb-rpc-mock-data
git checkout v0.118-rc1
pip install -r requirements.txt
pip install Werkzeug==2.2.2
python3 api/index.py > index.log 2>&1 &