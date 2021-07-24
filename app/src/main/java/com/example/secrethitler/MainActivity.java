package com.example.secrethitler;

import androidx.appcompat.app.AppCompatActivity;

import android.content.Context;
import android.content.Intent;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;

public class MainActivity extends AppCompatActivity {

    public Button createPrivateButton;
    public Button createPublicButton;
    public Button joinPrivateButton;
    public Button joinPublicButton;
    Context parentContext = this;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        createPrivateButton = findViewById(R.id.createPrivate);
        createPrivateButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(parentContext, CreatePrivateRoomActivity.class);
                startActivity(intent);
            }
        });

        createPublicButton = findViewById(R.id.createPublic);
        createPublicButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(parentContext, CreatePublicRoomActivity.class);
                startActivity(intent);
            }
        });

        joinPrivateButton = findViewById(R.id.joinPrivate);
        joinPrivateButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(parentContext, JoinPrivateRoomActivity.class);
                startActivity(intent);
            }
        });

        joinPublicButton = findViewById(R.id.joinPublic);
        joinPublicButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                Intent intent = new Intent(parentContext, JoinPublicRoomActivity.class);
                startActivity(intent);
            }
        });
    }
}
