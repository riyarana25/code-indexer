FROM python:3.13-slim

WORKDIR /analyzer

# Install dependencies
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt

# Copy the analyzer code
COPY . .

# Run the analyzer
CMD ["python", "analyzer.py"] 